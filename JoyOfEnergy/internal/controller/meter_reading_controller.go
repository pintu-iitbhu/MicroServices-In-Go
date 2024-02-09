package controller

import (
	"JoyOfEnergy/internal/constants"
	"JoyOfEnergy/internal/dto/request"
	"JoyOfEnergy/internal/errorhandling"
	"JoyOfEnergy/internal/service"
	"JoyOfEnergy/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type MeterReadingController struct {
	Validator           *validator.Validate
	MeterReadingService service.IMeterReadingService
}

func NewMeterReadingController(validator *validator.Validate, mrService service.IMeterReadingService) *MeterReadingController {
	return &MeterReadingController{
		Validator:           validator,
		MeterReadingService: mrService,
	}
}
func (m *MeterReadingController) PostMeterReading() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		log := logger.NewLogger()
		userId := c.Param(constants.UserId)
		meterId := c.Param(constants.MeterId)

		if err := m.Validator.Var(userId, constants.NotNilUUIDString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid userid"})
			return
		}
		if err := m.Validator.Var(meterId, constants.NotNilUUIDString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid meterId"})
			return
		}

		var req request.MeterReadingRequest
		if bindError := c.BindJSON(&req); bindError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request " + bindError.Error()})
			return
		}

		if validationErr := m.Validator.Struct(req); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "validation error " + validationErr.Error()})
		}

		serviceError := m.MeterReadingService.StoreMeterReading(userId, meterId, req)
		if serviceError != nil {
			log.Info("Meter reading service giving error", "StatusCode", serviceError.StatusCode)
			c.JSON(errorhandling.GetHttpStatus(serviceError.StatusCode), gin.H{})
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	return fn
}

func (m *MeterReadingController) GetMeterReading() gin.HandlerFunc {
	fn := func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{})
		return
	}
	return fn
}
