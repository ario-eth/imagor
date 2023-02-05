package main

import (
	"github.com/ario-eth/imagor/config"
	"github.com/ario-eth/imagor/config/awsconfig"
	"github.com/ario-eth/imagor/config/gcloudconfig"
	"github.com/ario-eth/imagor/config/vipsconfig"
	"os"
)

func main() {
	var server = config.CreateServer(
		os.Args[1:],
		vipsconfig.WithVips,
		awsconfig.WithAWS,
		gcloudconfig.WithGCloud,
	)
	if server != nil {
		server.Run()
	}
}
