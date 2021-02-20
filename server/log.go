package server

import (
	"fmt"
	"log"
)

const (
	logFormat        = "Method: %s Endpoint: %s Status: %d Params: %s"
	successLogFormat = logFormat + "\n"
	failedLogFormat  = logFormat + " Error: %s\n"
)

type loggingQuery struct {
	key   string
	value string
}

func parseParamsToString(params []loggingQuery) (paramsString string) {
	for _, param := range params {
		if len(paramsString) != 0 {
			paramsString += ","
		}

		paramsString += fmt.Sprintf("%s: %s", param.key, param.value)
	}

	return
}

func loggingSuccess(method string, endpoint string, status int, params []loggingQuery) {
	paramsString := parseParamsToString(params)
	log.Printf(successLogFormat, method, endpoint, status, paramsString)
}

func loggingFailed(method string, endpoint string, status int, params []loggingQuery, err error) {
	paramsString := parseParamsToString(params)
	log.Printf(failedLogFormat, method, endpoint, status, paramsString, err)
}
