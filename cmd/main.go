package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/dimaxdqwerty/go-web-app/pkg/api"
	"github.com/dimaxdqwerty/go-web-app/pkg/data"
	"github.com/dimaxdqwerty/go-web-app/pkg/db"

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

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalln(err)
	}
	r := mux.NewRouter()
	userData := data.NewUserData(conn)
	api.ServeUserResource(r, *userData)
	r.Use(mux.CORSMethodMiddleware(r))
	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalln(err)
	}
	err = http.Serve(listener, r)
	if err != nil {
		log.Fatalln(err)
	}

}
