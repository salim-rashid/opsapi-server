package openresty

import (
	"fmt"
	"net/http"
	"runtime"
)

func OpenrestyConf() {
	http.HandleFunc("/v1/openresty_install", OpenrestyInstall)

	// http.HandleFunc("/v1/openresty_uninstall", OpenrestyUninstall)

	// http.HandleFunc("/v1/openresty_update", OpenrestyUpdate)

	// http.HandleFunc("/v1/openresty_start", OpenrestyStart)
}

func OpenrestyInstall(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Do Openresty installation.")

	// Detect OS
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Detect Linux distribution

	// Install Openresty

	// Return a JSON for the installation status
}

// func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Do Openresty uninstallation.")
// }

// func OpenrestyUpdate(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Do Openresty update.")
// }

// func OpenrestyStart(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Do Openresty restart.")
// }
