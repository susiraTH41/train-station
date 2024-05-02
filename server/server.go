package server

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/susiraTH41/train-station/config"
	"github.com/susiraTH41/train-station/databases"
)

type echoServer struct {
	app 	*echo.Echo
	db 		databases.Database
	conf 	*config.Config
}


var (
	once 	sync.Once
	server 	*echoServer

)

func NewEchoServer(conf *config.Config,db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)



	once.Do(func() {
		server = &echoServer{
			app: 	echoApp,
			db: 	db,
			conf: 	conf,
		}	
	})
	return server
}


