package nginx

import (
	"fmt"
	"net/http"
)

// func PrintNginxPkg() {
// 	fmt.Println("Hello, Modules! This is Nginx package speaking v1.0.0!")
// }

func NginxConf() {
	http.HandleFunc("/v1/nginx_install", NginxInstall)

	http.HandleFunc("/v1/nginx_uninstall", NginxUninstall)

	http.HandleFunc("/v1/nginx_update", NginxUpdate)

	http.HandleFunc("/v1/nginx_start", NginxStart)
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
