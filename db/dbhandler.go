package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/norbux/gqlgen-todos/graph/model"
)

var db *sql.DB

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123"
	dbname = "go_test"
)

func Connect() error {
	var err error
	db, err = sql.Open(
		"postgres", 
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
			host,
			port,
			user,
			password,
			dbname,
		),
	)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}

func dbConnect() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
}

func GetCanarios() []*model.Canario {
	dbConnect()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM canarios")
	defer rows.Close()

	result := []*model.Canario{}

	for rows.Next() {
		canario := model.Canario{}
		rows.Scan(&canario.ID, &canario.Name, &canario.Email)
		result = append(result, &canario)
	}

	return result
}

func CreateCanario(c *model.Canario) *model.Canario {
	dbConnect()
	defer db.Close()

	res, _ := db.Query("INSERT INTO canarios (name, email) VALUES ($1, $2)", c.Name, c.Email)
	defer res.Close()

	canario := new(model.Canario)

	res.Scan(&canario.ID, &canario.Name, &canario.Email)

	return canario
}