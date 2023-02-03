package openresty

import (
	"fmt"
	"net/http"
)

func OpenrestyConf() {
	http.HandleFunc("/v1/openresty_install", OpenrestyInstall)

	http.HandleFunc("/v1/openresty_uninstall", OpenrestyUninstall)

	http.HandleFunc("/v1/openresty_update", OpenrestyUpdate)

	http.HandleFunc("/v1/openresty_start", OpenrestyStart)
}

func OpenrestyInstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Openresty installation.")
}

func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Openresty uninstallation.")
}

func OpenrestyUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Openresty update.")
}

func OpenrestyStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do Openresty restart.")
}
