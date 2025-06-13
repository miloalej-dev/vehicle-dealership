package internal

type Vehicle struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Plate        string  `json:"plate"`
	Color        string  `json:"color"`
	Year         int     `json:"year"`
	MaxSpeed     int     `json:"max_speed"`
	Passengers   int     `json:"passengers"`
	Fuel         string  `json:"fuel"`
	Transmission string  `json:"transmission"`
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}
