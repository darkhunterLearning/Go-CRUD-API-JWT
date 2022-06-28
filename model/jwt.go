package model

type Authentication struct {
	Email    string `json:"customer_email"`
	Password string `json:"customer_password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"customer_email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
