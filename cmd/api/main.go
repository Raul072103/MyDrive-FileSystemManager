package main

import (
	"MyDrive-FileSystemManager/internal/config"
	"MyDrive-FileSystemManager/internal/handlers"
	"MyDrive-FileSystemManager/internal/helpers"
	"MyDrive-FileSystemManager/internal/repository"
	"MyDrive-FileSystemManager/internal/repository/filerepo"
	"fmt"
	"log"
	"net/http"
	"os"
)

const portNumber = ":8080"
const inProduction = false

var app config.AppConfig

var fileRepo repository.FileManagerRepo = filerepo.New()

var infoLog *log.Logger
var errorLog *log.Logger

// main is the main entry point in the api
func main() {
	// TODO() Connect to DB

	repo := repository.New(&app, fileRepo)
	run(repo)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	err := server.ListenAndServe()
	log.Fatal(err)
}

func run(repo *repository.Repository) {
	setUpSessions()
	setUpAppConfig()

	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
}

// setUpAppConfig sets up the api configuration (app)
func setUpAppConfig() {
	app.InProduction = inProduction

	// TODO() Persist logs throughout server lifetime
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.InfoLog = infoLog
	app.ErrorLog = errorLog
}

func setUpSessions() {
	// TODO() Register any models for the session

	//gob.Register(models.Reservation{})
	//gob.Register(models.User{})
	//gob.Register(models.Room{})
	//gob.Register(models.Restriction{})
	//
	//session = scs.New()
	//session.Lifetime = 24 * time.Hour
	//session.Cookie.Persist = true
	//session.Cookie.SameSite = http.SameSiteLaxMode
	//session.Cookie.Secure = app.InProduction
}
