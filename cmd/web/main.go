package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/jerevick83/hotel_bookings/pkg/config"
	"github.com/jerevick83/hotel_bookings/pkg/handlers"
	"github.com/jerevick83/hotel_bookings/pkg/render"

	"log"
	"net/http"
	"time"
)

var port string = ":8080"

//app is assigned the application wide configuration struct
var app config.AppConfig
var session *scs.SessionManager

func main() {
	//Change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	templateCache, err := render.BaseLayoutCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	//app.templateCache is assigned the template cache function in the render package
	app.TemplateCache = templateCache
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)
	srv := &http.Server{
		Handler: routes(&app),
		Addr:    port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server running on port :8080")
	log.Fatal(srv.ListenAndServe())
}
