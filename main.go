package main

import (
	"github.com/ogios/simple-socket-server/server/normal"

	"github.com/CPTServ/TServ/addon/proxy"
	"github.com/CPTServ/TServ/addon/udps"
	"github.com/CPTServ/TServ/config"
	"github.com/CPTServ/TServ/log"
)

func main() {
	server, err := normal.NewSocketServer(config.GlobalConfig.Address)
	if err != nil {
		log.Error(nil, "Socket server error: &v", err)
		panic(err)
	}

	AddRouters(server)

	log.Info(nil, "Start serving")
	udps.StartUdps()
	proxy.StartProxy()
	if err := server.Serv(); err != nil {
		log.Error(nil, "Serv error: &v", err)
		panic(err)
	}
}
