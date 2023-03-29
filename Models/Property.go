package models

type Property struct {
	Id              int64
	Title           string
	Owner_id        string
	City            string
	Country         string
	Total_rooms     int
	Total_area      int
	Rating          float64
	Nei_details     string
	Price           int
	Avg_living_cost int
	Facilities      string
}
