package structs

type Data struct {
	Humidity    float64 `json:"humidity"`
	Pressure    int     `json:"pressure"`
	Tempreature float64 `json:"temperature"`
	CreatedOn   string  `json:"created_on"`
}
