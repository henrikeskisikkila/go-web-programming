package main

import (
	"go-web-programming/chitchat/data"
	"net/http"
)

//GET /err/?msg=
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		navbar := "public.navbar"
		if err == nil {
			navbar = "private.navbar"
		}
		generateHTML(writer, threads, "layout", navbar, "index")
	}
}
