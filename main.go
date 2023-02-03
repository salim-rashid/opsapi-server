package main

import (
	"homepage"
	"log"
	"openresty"
	"opsapi-server/nginx"
	"opsapi-server/varnish"
	"webserver"
)

func main() {
	// Calling homepage
	homepage.HomePageConf()

	// Calling Nginx configuration.
	nginx.NginxConf()

	// Calling Varnish configuration.
	varnish.VarnishConf()

	// Calling Openrest configuration.
	openresty.OpenrestyConf()

	errs := webserver.Run(":8080", ":10443", map[string]string{
		"cert": "./webserver/localhost.crt",
		"key":  "./webserver/localhost.key",
	})

	// This will run forever until channel receives error
	select {
	case err := <-errs:
		log.Printf("Error starting server (error: %s)", err)
	}
}
