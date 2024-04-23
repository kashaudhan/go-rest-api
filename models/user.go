package models

import (
	"errors"

	"rest.api/db"
	"rest.api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	password, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	isValid := utils.ComparePassword(u.Password, hashedPassword); 
	if !isValid {
		return errors.New("invalid credentials")
	}

	return nil
}
