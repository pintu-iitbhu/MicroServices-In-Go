package router

import (
	"JoyOfEnergy/internal/constants"
	"JoyOfEnergy/internal/controller"
	"JoyOfEnergy/internal/service"
	"JoyOfEnergy/internal/utility"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	healthController := controller.NewHealthController()
	requestValidator := utility.NewValidator()
	meterReadingService := service.NewMeterReadingService()
	meterReadingController := controller.NewMeterReadingController(requestValidator, meterReadingService)
	joyOfEnergy := router.Group(constants.ServiceName)
	{
		joyOfEnergy.GET(constants.HealthCheckPath, healthController.Status())
		joyOfEnergy.POST("store-reading/userId/{userId}/meterId/{meterId}", meterReadingController.PostMeterReading())
	}
	return router

}
