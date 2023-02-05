package main

import (
	"github.com/ario-eth/imagor"
	"github.com/ario-eth/imagor/imagorpath"
	"github.com/ario-eth/imagor/loader/httploader"
	"github.com/ario-eth/imagor/server"
	"github.com/ario-eth/imagor/storage/filestorage"
	"github.com/ario-eth/imagor/vips"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())

	// create and run imagor server programmatically
	server.New(
		imagor.New(
			imagor.WithLogger(logger),
			imagor.WithUnsafe(true),
			imagor.WithProcessors(vips.NewProcessor()),
			imagor.WithLoaders(httploader.New()),
			imagor.WithStorages(filestorage.New("./")),
			imagor.WithResultStorages(filestorage.New("./")),
			imagor.WithResultStoragePathStyle(imagorpath.SuffixResultStorageHasher),
		),
		server.WithPort(8000),
		server.WithLogger(logger),
	).Run()
}
