package service

import (
	"JoyOfEnergy/internal/dto/request"
	"JoyOfEnergy/internal/dto/response"
)

type IMeterReadingService interface {
	StoreMeterReading(userId string, meterId string, req request.MeterReadingRequest) *response.ErrorResponseDto
}

type MeterReadingService struct {
}

func NewMeterReadingService() *MeterReadingService {
	return &MeterReadingService{}
}

func (m *MeterReadingService) StoreMeterReading(userId string, meterId string, req request.MeterReadingRequest) *response.ErrorResponseDto {
	return &response.ErrorResponseDto{}
}
