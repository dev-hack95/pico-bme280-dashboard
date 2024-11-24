package handlers

import (
	"html/template"

	"github.com/dev-hack95/pico-bme280-dashboard/structs"
	"github.com/dev-hack95/pico-bme280-dashboard/utilities"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	templateDir := "templates"
	authtoken, _ := c.Cookie("auth_token")
	IsTokenValid, _ := utilities.VerifyToken(authtoken)
	if IsTokenValid {
		if c.GetHeader("HX-Request") == "true" {
			c.Header("HX-Redirect", "/dashboard")
			c.Status(200)
		} else {
			c.Redirect(302, "/dashboard")
		}
		return
	}

	templPath := "/login/login.html"
	tmpl, err := template.ParseFiles(templateDir + templPath)
	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		c.JSON(400, returnData)
		return
	}

	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		c.JSON(400, returnData)
		return
	}
}

func SignupForm(c *gin.Context) {
	returnData := utilities.ResponseJson{}
	flag := true
	templateDir := "templates"

	tmpl, err := template.ParseFiles(templateDir + "/login/signup.html")

	if err != nil {
		flag = false
		utilities.ErrorResponse(&returnData, err.Error())
		c.JSON(400, returnData)
		return
	}

	switch {
	case flag:
		err := tmpl.Execute(c.Writer, nil)
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			return
		}
		return
	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}
}

func Dashboard(c *gin.Context) {
	username, _ := utilities.GetUserName(c)
	returnData := utilities.ResponseJson{}
	flag := true
	templateDir := "templates"

	tmpl, err := template.ParseFiles(templateDir + "/dashboard.html")

	if err != nil {
		flag = false
		utilities.ErrorResponse(&returnData, err.Error())
		c.JSON(400, returnData)
		return
	}

	switch {
	case flag:
		data := structs.User{Username: username}
		err := tmpl.Execute(c.Writer, data)
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			return
		}
		return
	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}
}
