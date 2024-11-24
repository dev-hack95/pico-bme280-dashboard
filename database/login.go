package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func UserLogin(username string) (string, string, error) {
	var user string
	var pass string

	db, err := sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		return "", "", err
	}

	defer db.Close()

	query := "SELECT user_name, password FROM users where user_name = ?;"

	rows, err := db.Query(query, username)

	if err != nil {
		return "", "", err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user,
			&pass,
		)

		if err != nil {
			return "", "", err
		}
	}

	return user, pass, nil

}

func GetPasswordByUserName(username string) (string, error) {
	var pass string

	db, err := sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		return "", err
	}

	defer db.Close()

	query := "SELECT password FROM users where user_name = ?;"

	rows, err := db.Query(query, username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&pass,
		)

		if err != nil {
			return "", err
		}
	}

	return pass, nil

}

func CreateUser(username, password string) bool {

	db, err := sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		return false
	}

	defer db.Close()

	query := "INSERT INTO users (user_name, password) VALUES (?, ?);"

	result, err := db.Exec(query, username, password)

	if err != nil {
		return false
	}

	RowsAffected, _ := result.RowsAffected()

	return RowsAffected > 0
}

func CheckUserIsPresentInB(username string) bool {
	var data int
	db, err := sql.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		return false
	}

	defer db.Close()

	query := "SELECT count(user_name) FROM users where user_name = ?;"

	rows, err := db.Query(query, username)

	if err != nil {
		return false
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&data,
		)

		if err != nil {
			return false
		}
	}

	if data == 0 {
		return false
	} else {
		return true
	}

}
