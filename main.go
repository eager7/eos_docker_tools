package main

import (
	"github.com/eager7/eos_docker_tools/http_server"
	"github.com/eager7/go/mlog"
)

var log = mlog.L

func main() {
	log.Debug("this is a cdt tool server for eos smart contract")
	if err := http_server.InitializeGin(); err != nil {
		log.Panic(err)
	}
}
