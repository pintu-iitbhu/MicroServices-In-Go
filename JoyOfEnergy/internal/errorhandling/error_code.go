package errorhandling

import "net/http"

type StatusCode string

var httpStatusMap map[StatusCode]int
var errorMessageMap map[StatusCode]string

const (
	JOY500 StatusCode = "JOY500"
	JOY400 StatusCode = "JOY400"
)

func NewHttpStatusMap() {
	httpStatusMap = make(map[StatusCode]int)
	httpStatusMap[JOY500] = http.StatusInternalServerError
	httpStatusMap[JOY400] = http.StatusBadRequest
}

func NewErrorStatusMessage() {
	errorMessageMap = make(map[StatusCode]string)

	errorMessageMap[JOY500] = "Inetrnal server error"
	errorMessageMap[JOY400] = "Bad request"
}

func GetHttpStatus(statusCode StatusCode) int {
	if httpStatus, ok := httpStatusMap[statusCode]; ok {
		return httpStatus
	}
	return http.StatusInternalServerError
}

func GetErrorMessage(statusCode StatusCode) string {
	if errorMessage, ok := errorMessageMap[statusCode]; ok {
		return errorMessage
	}
	return "An error has occurred"
}
