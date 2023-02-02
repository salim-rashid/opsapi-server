package varnish

import (
	"fmt"
	"net/http"
)

func PrintVarnishPkg() {
	fmt.Println("Hello, Modules! This is package varnish speaking v1.0.0!")
}

func VarnishConf() {
	http.HandleFunc("/v1/varnish_install", VarnishInstall)

	http.HandleFunc("/v1/varnish_uninstall", VarnishUninstall)

	http.HandleFunc("/v1/varnish_update", VarnishUpdate)

	http.HandleFunc("/v1/varnish_start", VarnishStart)
}

func VarnishInstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Varnish installation.")
}

func VarnishUninstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Varnish uninstallation.")
}

func VarnishUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Varnish update.")
}

func VarnishStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Varnish restart.")
}
