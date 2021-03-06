package main

import (
	"net/http"
	"time"
)

func main() {
	p(config)
	p("ChicChat", version(), "started at", config.Address)

	//Handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	//route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)

	//Starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
