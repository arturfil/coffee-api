package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/coffee-api/db"
	"github.com/arturfil/coffee-api/router"
	"github.com/arturfil/coffee-api/services"
)

type Config struct {
    Port string
}

type Application struct {
    Config Config
    Models services.Models
}

// global port varialbe for both main and serve
var port = os.Getenv("PORT")

func (app *Application) Serve() error {
    fmt.Println("API listening on port", port)

    srv := &http.Server{
        Addr: fmt.Sprintf(":%s", port),
        Handler: router.Routes(),
    }

    return srv.ListenAndServe()
}

func main() {
    var cfg Config
    cfg.Port = port

    dsn := os.Getenv("DSN")
    dbConn, err := db.ConnectPostgres(dsn)
    if err != nil {
        log.Fatal("Cannot connect to database", err)
    }

    defer dbConn.DB.Close()

    app := &Application {
        Config: cfg,
        Models: services.New(dbConn.DB),
    }

    err = app.Serve()
    if err != nil {
        log.Fatal(err)
    }
}






