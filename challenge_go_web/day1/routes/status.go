package routes

import (
	"net/http" 
) 

func GetHttpStatus(statusCode int) interface{}{
	status := map[string]interface{}{
        "httpStatusCode": statusCode,
        "message": http.StatusText(statusCode),
    }

	return status
}