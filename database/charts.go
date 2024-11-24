package database

import (
	"database/sql"
	"os"
	"strings"

	"github.com/dev-hack95/pico-bme280-dashboard/structs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cast"
)

func GetChartDetails() ([]structs.Data, error) {
	var data []structs.Data
	var humidity, tempreate, pressure string

	db, err := sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT humidity, pressure, temperature, timestamp FROM sensor_readings ORDER BY timestamp DESC LIMIT 1000"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var record structs.Data
		err := rows.Scan(
			&humidity,
			&pressure,
			&tempreate,
			&record.CreatedOn,
		)

		if err != nil {
			return nil, err
		}

		record.Humidity = cast.ToFloat64(humidity)
		record.Pressure = cast.ToInt(pressure)
		record.Tempreature = cast.ToFloat64(removeLastC(tempreate))

		data = append(data, record)

	}

	return data, nil
}

func removeLastC(input string) string {
	if strings.HasSuffix(input, "C") {
		return input[:len(input)-1]
	}
	return input
}
