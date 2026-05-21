package entity

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDetail struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	FullName        string `json:"full_name"`
	Phone           string `json:"phone"`
	ShippingAddress string `json:"shipping_address"`
}
