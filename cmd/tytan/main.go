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

	defer db.Close()

	router := router.NewRouter(
		validator.New(),
		db,
	)

	http.ListenAndServe(":8080", router)
	//http.ListenAndServe(":"+config.Server.Port, router)
}
