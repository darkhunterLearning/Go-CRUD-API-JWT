package model

type Customer struct {
	ID                uint32 `json:"id"`
	Unique_Id         string `json:"unique_id"`
	Customer_Name     string `json:"customer_name"`
	Customer_Phone    string `json:"customer_phone"`
	Customer_Address  string `json:"customer_address"`
	Customer_Email    string `json:"customer_email"`
	Customer_Password string `json:"customer_password"`
	Customer_Role     string `json:"customer_role"`
}
