package entity

import "time"

type Order struct {
	ID        int
	UserID    int
	OrderDate time.Time
	Status    string
}
