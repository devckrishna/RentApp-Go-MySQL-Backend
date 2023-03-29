package models

type Review struct {
	Id          string
	Property_id int
	User_id     int
	Body        string
	Rating      float64
}
