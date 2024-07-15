package main

import (
	"github.com/susiraTH41/train-station/config"
	"github.com/susiraTH41/train-station/databases"
	"github.com/susiraTH41/train-station/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)
	server.Start()
}
