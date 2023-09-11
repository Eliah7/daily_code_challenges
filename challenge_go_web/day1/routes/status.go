package routes

import "net/http" 

type Status struct {
	httpStatusCode int
	message string
}

func GetHttpStatus(statusCode int) Status{
	return Status{httpStatusCode: statusCode, message: http.StatusText(statusCode)}
}