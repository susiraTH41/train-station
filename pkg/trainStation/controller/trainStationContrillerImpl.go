package controller

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/susiraTH41/train-station/pkg/custom"

	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
	_trainStationService "github.com/susiraTH41/train-station/pkg/trainStation/service"
)

type trainStationControllerImpl struct {
	trainStationService _trainStationService.TrainStationService
}

func NewTrainStationControllerImpl( 
	trainStationService _trainStationService.TrainStationService ,
) TrainStationController {
	return &trainStationControllerImpl{trainStationService}
}


func (c *trainStationControllerImpl) GetStationNearMe(pctx echo.Context) error{
	stationFilter := new(_trainStationModel.StationFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(stationFilter); err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}


	stationModelList , err := c.trainStationService.GetStationNearMe(stationFilter)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, stationModelList)

}

func (c *trainStationControllerImpl) GetStationNearMeOnPage(pctx echo.Context) error{
	stationFilter := new(_trainStationModel.StationPaginateFilter)

	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(stationFilter); err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}


	stationModelList , err := c.trainStationService.GetStationNearMeOnPage(stationFilter)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}


	return pctx.JSON(http.StatusOK, stationModelList)
}


