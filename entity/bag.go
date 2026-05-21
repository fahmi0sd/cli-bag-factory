package entity

import "time"

type Bag struct {
	ID         int
	CategoryID int
	Name       string
	Material   string
	Price      int
	Stock      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
