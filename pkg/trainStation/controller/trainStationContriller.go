package controller

import "github.com/labstack/echo/v4"



type TrainStationController interface {
	GetStationNearMe(pctx echo.Context) error
	GetStationNearMeOnPage(pctx echo.Context) error


}