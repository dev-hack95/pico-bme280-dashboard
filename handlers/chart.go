package handlers

import (
	"html/template"

	"github.com/dev-hack95/pico-bme280-dashboard/controllers"
	"github.com/dev-hack95/pico-bme280-dashboard/utilities"
	"github.com/gin-gonic/gin"
)

func GetHumidityChartDetails(c *gin.Context) {
	var err error
	returnData := utilities.ResponseJson{}
	errAuth := utilities.GetUserSessionDetails(c)
	flag := true

	if errAuth != nil {
		flag = false
		c.Header("Hx-Redirect", "/")
		c.Status(200)
		return
	}

	switch {
	case flag:
		returnData, err = controllers.GetChartDetails()
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			c.Status(400)
			return
		}

		if returnData.Msg == "Success" {
			tmpl := `
			<div id="outerSwapHumidity"
				 hx-get="/pico/dashboard/chart/humidity"
				 hx-trigger="every 3s"
				 hx-swap="outerHTML"
				 class="chart-container">
				<h2>Humidity</h2>
				<div id="outerChartHumidity" class="chart"></div>
				<script>
					var humidityMetric = 'humidity'; // Declare the variable
					initChart('outerChartHumidity', {{.}}, humidityMetric); // Pass it to initChart
				</script>
			</div>
			`
			t := template.Must(template.New("outer").Parse(tmpl))
			t.Execute(c.Writer, returnData.Data)
		}

	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}

}

func GetTempreatureChartDetails(c *gin.Context) {
	var err error
	returnData := utilities.ResponseJson{}
	errAuth := utilities.GetUserSessionDetails(c)
	flag := true

	if errAuth != nil {
		flag = false
		c.Header("Hx-Redirect", "/")
		c.Status(200)
		return
	}

	switch {
	case flag:
		returnData, err = controllers.GetChartDetails()
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			c.Status(400)
			return
		}

		if returnData.Msg == "Success" {
			tmpl := `
			<div id="outerSwapTemperature"
				 hx-get="/pico/dashboard/chart/temperature"
				 hx-trigger="every 3s"
				 hx-swap="outerHTML"
				 class="chart-container">
				<h2>Temperature</h2>
				<div id="outerChartTemperature" class="chart"></div>
				<script>
					var temperatureMetric = 'temperature'; // Declare the variable
					initChart('outerChartTemperature', {{.}}, temperatureMetric); // Pass it to initChart
				</script>
			</div>
			`
			t := template.Must(template.New("outer").Parse(tmpl))
			t.Execute(c.Writer, returnData.Data)
		}

	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}

}

func GetPressureChartDetails(c *gin.Context) {
	var err error
	returnData := utilities.ResponseJson{}
	errAuth := utilities.GetUserSessionDetails(c)
	flag := true

	if errAuth != nil {
		flag = false
		c.Header("Hx-Redirect", "/")
		c.Status(200)
		return
	}

	switch {
	case flag:
		returnData, err = controllers.GetChartDetails()
		if err != nil {
			utilities.ErrorResponse(&returnData, err.Error())
			c.JSON(400, returnData)
			c.Status(400)
			return
		}

		if returnData.Msg == "Success" {
			tmpl := `
			<div id="outerSwapPressure"
				 hx-get="/pico/dashboard/chart/pressure"
				 hx-trigger="every 3s"
				 hx-swap="outerHTML"
				 class="chart-container">
				<h2>Pressure</h2>
				<div id="outerChartPressure" class="chart"></div>
				<script>
					var pressureMetric = 'pressure'; // Declare the variable
					initChart('outerChartPressure', {{.}}, pressureMetric); // Pass it to initChart
				</script>
			</div>
			`
			t := template.Must(template.New("outer").Parse(tmpl))
			t.Execute(c.Writer, returnData.Data)
		}

	default:
		utilities.ErrorResponse(&returnData, "Something went wrong")
		c.JSON(400, returnData)
		return
	}
}
