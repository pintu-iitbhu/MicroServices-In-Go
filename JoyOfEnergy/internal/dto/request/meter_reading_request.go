package request

type MeterReadingRequest struct {
	Reading float64 `json:"reading" validate:""`
}
