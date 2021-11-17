package models

type Customer struct {
	Id       int    `json:"id,string"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}
