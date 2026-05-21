package handler

import (
	"database/sql"
	"errors"
	"log"

	"github.com/fahmi0sd/cli-bag-factory/entity"
)

// function login
func (h *CLIHandler) Login(email, password string) (*entity.User, error) {
	query := `SELECT id, email, role FROM users WHERE email = ? AND password = ?`

	var user entity.User
	err := h.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("email atau password salah")
		}
		log.Println("Database error:", err)
		return nil, err
	}

	return &user, nil
}

// function register customer
func (h *CLIHandler) RegisterCustomer(email, password, fullName, phone, address string) error {
	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}

	// Insert data user
	queryUser := `INSERT INTO users (email, password, role) VALUES (?, ?, 'Customer')`
	res, err := tx.Exec(queryUser, email, password)
	if err != nil {
		tx.Rollback()
		return err
	}

	// check whether the user ID already exists or not ?
	userID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert to table user_details
	queryDetail := `INSERT INTO user_details (user_id, full_name, phone, shipping_address) VALUES (?, ?, ?, ?)`
	_, err = tx.Exec(queryDetail, userID, fullName, phone, address)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
