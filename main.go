package main

import (
	"fmt"
	"jstcloud-golang/db"
	"jstcloud-golang/service"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.HelloWorldHandler)
	http.HandleFunc("/counter", service.CounterHandler)
	http.HandleFunc("/counter/increment", service.CounterHandler)
	http.HandleFunc("/counter/decrement", service.CounterHandler)

	port := os.Getenv("SERVER_PORT") // 从环境变量中获取端口号
	if port == "" {
		port = "8080" // 如果未设置，则使用默认端口号8080
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
