package main

import (
	"database/sql"
	"go/main/api/router"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "file:database.db?mode=memory&cache=shared")
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

	r := router.NewRouter(db)
	http.ListenAndServe(":8080", r)
}
