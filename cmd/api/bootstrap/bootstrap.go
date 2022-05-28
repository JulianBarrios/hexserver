package bootstrap

import (
	"github.com/JulianBarrios/hexserver/infrastructure/server"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
