package handlers

import (
	"github.com/dev-hack95/pico-bme280-dashboard/controllers"
	"github.com/dev-hack95/pico-bme280-dashboard/utilities"
	"github.com/gin-gonic/gin"
)

func CreateUserAccount(c *gin.Context) {
	var err error
	returnData := utilities.ResponseJson{}
	flag := true

	username := c.Request.FormValue("user_name")
	password := c.Request.FormValue("password")
	confirm_password := c.Request.FormValue("confirm_password")

	if password != confirm_password {
		utilities.ErrorResponse(&returnData, "Password and confirm password do not match")
		c.JSON(400, returnData)
		return
	}

	switch {
	case !utilities.IsEmpty(username) && !utilities.IsEmpty(password) && flag:
		returnData, err = controllers.CreateUser(username, password)
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			return
		}

		if returnData.Msg == "Success" {
			c.Header("HX-Redirect", "/")
			c.Status(200)
			return
		}

	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}
}

func UserLogin(c *gin.Context) {
	var err error
	returnData := utilities.ResponseJson{}
	flag := true

	username := c.Request.FormValue("user_name")
	password := c.Request.FormValue("password")

	switch {
	case !utilities.IsEmpty(username) && !utilities.IsEmpty(password) && flag:
		returnData, err = controllers.UserLogin(username, password)
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			return
		}

		if returnData.Msg == "Success" && returnData.Data != "" {
			c.SetCookie("auth_token", returnData.Data.(string), 3600, "/", "", false, true)
			c.Header("HX-Redirect", "/dashboard")
			c.Status(200)
			return
		}
	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}
}

func UserLogout(c *gin.Context) {
	errAuth := utilities.GetUserSessionDetails(c)

	if errAuth != nil {
		c.JSON(400, gin.H{"error": "<div class='error-message'>" + errAuth.Error() + "</div>"})
		return
	}
	// unset cookie
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.Header("HX-Redirect", "/")
	c.Status(200)
}
