package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/markledger/bookings/internal/config"
	"github.com/markledger/bookings/internal/driver"
	"github.com/markledger/bookings/internal/handlers"
	"github.com/markledger/bookings/internal/models"
	"github.com/markledger/bookings/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run() (*driver.DB, error) {
	//What am I going to put in session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost password= user=mark.ledger database=bookings port=5432")

	tc, err := render.CreateTemplateCache()
	if err != nil {

		log.Fatal(err)
		log.Fatal("Cannot create template cache")
		return db, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	return db, nil
}
