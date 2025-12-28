package main

import (
	"database/sql"
	"net/http"
	"tytan-api/internal/router"
	"tytan-api/internal/util/validator"

	_ "modernc.org/sqlite"
)

func main() {
	// config := config.LoadConfig()

	db, err := sql.Open("sqlite", "file:database.db?cache=shared")
	// db, err := sql.Open(config.Database.Driver, "file:"+config.Database.Host+"?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY, 
			
			username TEXT NOT NULL,
			email TEXT NOT NULL, 
			password TEXT NOT NULL, 
			created_at DATETIME NOT NULL DEFAULT (datetime('now')),
			updated_at DATETIME NOT NULL DEFAULT (datetime('now')),

			UNIQUE (id),
			UNIQUE (username),
			UNIQUE (email)
		);

		CREATE INDEX IF NOT EXISTS idx_credentials_username ON users (username, password);
		CREATE INDEX IF NOT EXISTS idx_credentials_email ON users (email, password);
		CREATE INDEX IF NOT EXISTS idx_id ON users (id);
	`)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	router := router.NewRouter(
		validator.New(),
		db,
	)

	http.ListenAndServe(":8080", router)
	//http.ListenAndServe(":"+config.Server.Port, router)
}
