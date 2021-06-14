package main

import (
	"go-web-programming/chitchat/data"
	"net/http"
)

//Get /login
//Show login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

//POST /authenticate
//Authentice the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
