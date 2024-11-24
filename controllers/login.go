package controllers

import (
	"errors"

	"github.com/dev-hack95/pico-bme280-dashboard/database"
	"github.com/dev-hack95/pico-bme280-dashboard/utilities"
)

func UserLogin(username, password string) (returnData utilities.ResponseJson, err error) {
	userpassword, _ := database.GetPasswordByUserName(username)

	if !utilities.VerifyPassword(userpassword, password) {
		return returnData, errors.New("password is wrong")
	}

	token, err := utilities.CreateToken(username, true)

	if err != nil {
		return returnData, err
	}

	utilities.NewSuccessResponse(&returnData, token)
	return
}

func CreateUser(username, password string) (returnData utilities.ResponseJson, err error) {

	if database.CheckUserIsPresentInB(username) {
		return returnData, errors.New("user already present in db")
	}

	hashpassword, err := utilities.HashPassword(password)

	if err != nil {
		return returnData, err
	}

	IsUserCreated := database.CreateUser(username, hashpassword)

	if !IsUserCreated {
		return returnData, errors.New("error occured while creating new user")
	}

	utilities.SuccessResponse(&returnData)
	return
}
