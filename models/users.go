package models

import (
	"errors"

	"github.com/atharv-g-kulkarni/go_rest_api/db"
	"github.com/atharv-g-kulkarni/go_rest_api/utils"
)

type User struct {
	ID       int64
	EMAIL    string `binding:"required"`
	PASSWORD string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.PASSWORD)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.EMAIL, hashedPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.ID = userID
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email=?`
	row := db.DB.QueryRow(query, u.EMAIL)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials error")
	}
	passwordIsValid := utils.CheckPasswordHash(u.PASSWORD, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials error")
	}
	return nil
}
