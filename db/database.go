package db

import (
	"database/sql"
	"os"
)

var database *sql.DB

func InitDatabase() {
	db, err := sql.Open("mysql", os.Getenv("DB_STR"))

	if err != nil {
		panic(err.Error())
	}

	database = db
}

func AddNew(n string) (int64, error) {
	res, err := database.Exec("INSERT INTO urls (original_url) VALUES (?)", n)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetOriginal(id int) (string, error) {
	var original string

	err := database.QueryRow("SELECT original_url FROM urls WHERE id = ?", id).Scan(&original)

	if err != nil {
		return "", err
	}

	return original, nil
}
