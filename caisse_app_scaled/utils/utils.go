package utils

import "caisse-app-scaled/caisse_app_scaled/logger"

var GATEWAY string = "172.17.0.1" //172.17.0.1
var API_MERE string = "http://" + GATEWAY + ":8090"
var API_LOGISTIC string = "http://" + GATEWAY + ":8091"

func Errnotnil(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
