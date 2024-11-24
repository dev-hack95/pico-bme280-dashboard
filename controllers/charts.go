package controllers

import (
	"github.com/dev-hack95/pico-bme280-dashboard/database"
	"github.com/dev-hack95/pico-bme280-dashboard/utilities"
)

func GetChartDetails() (returnData utilities.ResponseJson, err error) {
	data, err := database.GetChartDetails()

	if err != nil {
		return returnData, err
	}

	utilities.NewSuccessResponse(&returnData, data)
	return
}
