package models

type User struct {
	Id             int64
	Name           string
	Email          string
	Phone          *string
	Password       string
	Age            *int
	Gender         *bool
	Marital_status *bool
	Photo          *string
	Is_host        *bool
}
