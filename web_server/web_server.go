package web_server

import (
	"fmt"
	"log"
	"net/http"
)

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

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Edgeone API Server - https://edgeone.cloud/api")
}

func NginxInstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Nginx installation.")
}

func NginxUninstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Nginx uninstallation.")
}

func NginxUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Nginx update.")
}

func NginxStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Nginx restart.")
}
