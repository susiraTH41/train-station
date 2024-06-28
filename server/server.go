package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func (s *echoServer) Start(){
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.TimeOut) 


	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeOutMiddleware)

	s.app.GET("/v1/auth",s.AuthRequir)

	s.initTrainStationRouter(s.authRequired())

	s.app.GET("/v1/health" , s.healthCheck)

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitChan)
	
	s.httpListener()
}

func (s *echoServer) httpListener(){
	url := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s" , err.Error())
	}
}

func (s *echoServer) gracefullyShutdown(quitChan chan os.Signal){
	ctx := context.Background()

	<-quitChan
	s.app.Logger.Info("Shutting down server...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s" , err.Error())
	}
}


func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}



func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:middleware.DefaultSkipper,
		ErrorMessage: "Request timeout",
		Timeout: timeout * time.Second,
	})
}


func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET,echo.POST },
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}, 

	})
}

func getBodyLimitMiddleware(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}

func (s *echoServer) AuthRequir(pctx echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed",
		})
	}

	// Set cookie
	pctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   t,
		Expires: time.Now().Add(time.Hour * 72),
		
	})

	return pctx.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})

}




func (s *echoServer) authRequired() echo.MiddlewareFunc  {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(pctx echo.Context) error {
			cookie, err := pctx.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}
			
			jwtSecretKey := []byte(os.Getenv("SECRET"))
		
			token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtSecretKey, nil
			})
		  
			if err != nil || !token.Valid {
				return err
			}
		  
			return next(pctx)
		}

	}

}