package main

import (
	"REST-APIs-Go/handlers"
	"fmt"
	"net/http"
	"os"
)


func main(){
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	err := http.ListenAndServe("localhost:11111", nil)

	if(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}
}