package main

import (
	"./controllers"
	"./router"
	"fmt"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr:         ":8013",
		Handler:      router.GetRouter(),
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	fmt.Println("server is starting")

	err := server.ListenAndServe()
	go controllers.Echo()
	if err!=nil{
		panic(err)
	}
}
