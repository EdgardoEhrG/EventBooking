package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"
)

type User struct {
	ID 			int64
	Email 		string 	`binding:"required"`
	Password 	string 	`binding:"required"`
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = ?`

	row := db.DB.QueryRow(query, id)

	var user User

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	smtm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer smtm.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := smtm.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()

	u.ID = userId

	return err
}

func (user User) Delete() error {
	query := `DELETE FROM users WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&user.ID)

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}