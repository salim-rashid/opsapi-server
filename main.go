package main

import (
	"log"
	"net/http"
	"web_server"
	// "opsapi/nginx"
)

func main() {

	http.HandleFunc("/", web_server.HomePage)

	errs := web_server.Run(":8080", ":10443", map[string]string{
		"cert": "./web_server/localhost.crt",
		"key":  "./web_server/localhost.key",
	})

	// This will run forever until channel receives error
	select {
	case err := <-errs:
		log.Printf("Error starting server (error: %s)", err)
	}
}
