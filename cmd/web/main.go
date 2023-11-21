package main

import (
	"log"
	"os"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
}

type server struct {
	host string
	port string
	url  string
}

func main() {
	server := server{
		host: "localhost",
		port: "25567",
		url:  "http://localhost:25567",
	}

	app := &application{
		appName: "Tinify",
		server:  server,
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Lshortfile),
	}

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}
}
