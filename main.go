package main

import (
	"fmt"
	"log"
	"net/http"
	"opsapi/filemanager"
	"opsapi/nginx"
	"opsapi/pop"
	"opsapi/varnish"
)

func test() {

	fmt.Println("Hello, Modules OPSAPI filemanager package!")

	filemanager.PrintFileManager()

	nginx.PrintNginxPkg()

	varnish.PrintVarnishPkg()

	pop.PrintPopPkg()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Edgeone API Server - https://edgeone.cloud/api")
}

func generateCerts() {
	// generate a self-signed certificate
	//cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
}

func Run(addr string, sslAddr string, ssl map[string]string) chan error {

	errs := make(chan error)

	// Starting HTTP server
	go func() {
		log.Printf("Staring HTTP service on %s ...", addr)

		if err := http.ListenAndServe(addr, nil); err != nil {
			errs <- err
		}

	}()

	// Starting HTTPS server
	go func() {
		log.Printf("Staring HTTPS service on %s ...", sslAddr)
		if err := http.ListenAndServeTLS(sslAddr, ssl["cert"], ssl["key"], nil); err != nil {
			errs <- err
		}
	}()

	return errs
}

func main() {

	http.HandleFunc("/", homePage)

	errs := Run(":8080", ":10443", map[string]string{
		"cert": "localhost.crt",
		"key":  "localhost.key",
	})

	// This will run forever until channel receives error
	select {
	case err := <-errs:
		log.Printf("Error starting server (error: %s)", err)
	}

}
