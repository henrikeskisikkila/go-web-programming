package main

import (
	"fmt"
	"net/http"

	"chitchat/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, data.Hello())
	/*
		fmt.Fprint(writer, "Hello")
		threads, err := data.Threads()
		if err != nil {
			error_message(writer, request, "Cannot get threads")
		} else {
			fmt.Fprint(writer, threads)
		}
	*/
}
