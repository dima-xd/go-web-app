package main

import (
	"go-web-app/pkg/controller"
	"go-web-app/pkg/data"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

	"go-web-app/pkg/db"

	"github.com/gorilla/mux"
)

var (
	serverPort = os.Getenv("SERVER_PORT")
	host       = os.Getenv("DB_USERS_HOST")
	port       = os.Getenv("DB_USERS_PORT")
	user       = os.Getenv("DB_USERS_USER")
	dbname     = os.Getenv("DB_USERS_DBNAME")
	password   = os.Getenv("DB_USERS_PASSWORD")
	sslmode    = os.Getenv("DB_USERS_SSL")
)

func init() {
	if serverPort == "" {
		serverPort = ":8081"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "web-app"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

var templates *template.Template

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalln(err)
	}
	templates = template.Must(template.ParseGlob("template/*.html"))
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	userData := data.NewUserData(conn)
	controller.ServeLoginResource(r, *userData, templates)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalln(err)
	}
	err = http.Serve(listener, r)
	if err != nil {
		log.Fatalln(err)
	}
}
