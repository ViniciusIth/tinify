package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) listenAndServer() error {
	srv := http.Server{
		Handler:     a.routes(),
		Addr:        fmt.Sprintf("%s:%s", a.server.host, a.server.port),
		ReadTimeout: 300 * time.Second,
	}

	a.infoLog.Printf("Server listening on: %s\n", a.server.url)

	return srv.ListenAndServe()
}
