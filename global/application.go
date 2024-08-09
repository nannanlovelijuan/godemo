package global

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"gitlab.ezrpro.in/godemo/etc"
)

type Name string
type Version string

type Application struct {
	Info   *ApplicationInfo
	Config *etc.Config
	quit   chan os.Signal
	wait   *sync.WaitGroup
}

type ApplicationInfo struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Env      string `json:"env"`
	Host     string `json:"host"`
	Hostname string `json:"hostname"`
	RootDir  string `json:"root_dir"`
}

func NewApplication(name string, version string) *Application {
	//获取应用主机IP
	host, err := getHost()
	if err != nil {
		log.Fatal(err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	env := os.Getenv("ENV")
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	rootDir := filepath.Dir(ex)

	info := ApplicationInfo{
		Name:     string(name),
		Version:  string(version),
		Env:      env,
		Host:     host,
		Hostname: hostName,
		RootDir:  rootDir,
	}

	quit := make(chan os.Signal, 5)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	wait := &sync.WaitGroup{}

	config, err := etc.NewConfig(name)
	application := &Application{
		Info:   &info,
		Config: config,
		quit:   quit,
		wait:   wait,
	}

	log.Printf("Application Info: %+v", application)
	return application
}

func (a *Application) Run(start func(a *Application)) {
	a.wait.Add(1)

	go func() {
		fmt.Println("Application start wait quit signal...")
		<-a.quit
		a.wait.Done()
	}()
	start(a)
}

func (a *Application) Wait() {
	fmt.Println("Application wait ...")
	a.wait.Wait()
}

func getHost() (host string, err error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("Failed to get interface addrs:", addrs)
		host = ""
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
				host = ipnet.IP.String()
				return
			}
		}
	}
	return "", fmt.Errorf("failed to get host")
}
