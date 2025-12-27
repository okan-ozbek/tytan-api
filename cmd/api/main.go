package main

import (
	"database/sql"
	"net/http"
	"tytan-api/api/router"
	"tytan-api/config"
	"tytan-api/util/validator"

	_ "modernc.org/sqlite"
)

func main() {
	config := config.LoadConfig()

	db, err := sql.Open(config.Database.Driver, "file:"+config.Database.Host+"?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, username TEXT, password TEXT, created_at DATETIME);`)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := router.NewRouter(
		validator.New(),
		db,
	)
	http.ListenAndServe(":8080", router)
}
