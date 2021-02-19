package server

import "log"

const logFormat = "Method: %s Endpoint: %s Status: %d Description: %s\n"

func loggingSuccess(method string, endpoint string, status int) {
	log.Printf(logFormat, method, endpoint, status, "success")
}
