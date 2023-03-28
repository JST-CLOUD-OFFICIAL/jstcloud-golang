package main

import (
	"fmt"
	"jstcloud-golang/db"
	"jstcloud-golang/service"
	"log"
	"net/http"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/counter", service.CounterHandler)
	http.HandleFunc("/counter/increment", service.CounterHandler)
	http.HandleFunc("/counter/decrement", service.CounterHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
