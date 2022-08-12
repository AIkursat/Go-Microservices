package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {}

func main() {

	app := Config{}

	log.Println("Starting broker service on port", webPort)

	// define htttp server
	srv := &http.Server{
		Addr : fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),

	}
    // start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
  
}