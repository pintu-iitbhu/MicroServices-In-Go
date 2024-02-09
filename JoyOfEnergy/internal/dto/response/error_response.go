package response

import "JoyOfEnergy/internal/errorhandling"

type ErrorResponseDto struct {
	StatusCode errorhandling.StatusCode
}
