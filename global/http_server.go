package global

import (
	"net/http"
	"strconv"
)

func NewHttpServer(port int) *http.Server {

	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: NewGinEngine(),
	}
}
