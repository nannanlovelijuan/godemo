package etc

import (
	"fmt"

	remote "gitlab.ezrpro.in/arch/agollo/viper-remote"
	"gitlab.ezrpro.in/arch/config"
)

type Config struct {
	AppId           string
	conf            *config.Config
	apolloServerUrl string
}

func NewConfig(appId string) (*Config, error) {
	remote.SetAppID(appId)

	conf := config.NewConfig()
	conf.AddConfigPath(".")
	conf.SetConfigName("config")
	conf.SetConfigType("json")

	err := conf.Parse()
	if err != nil {
		fmt.Printf("Parse config error%v\n", err)
		return nil, err
	}

	serverUrl := conf.GetString("apollo.serverUrl")

	err = conf.AddRemoteProvider("apollo", serverUrl, "application")
	if err != nil {
		fmt.Printf("AddRemoteProvider error%v\n", err)
		return nil, err
	}
	conf.Parse()

	return &Config{
		AppId:           appId,
		conf:            conf,
		apolloServerUrl: serverUrl,
	}, nil
}
