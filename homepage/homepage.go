package homepage

import (
	"fmt"
	"net/http"
)

func HomePageConf() {
	http.HandleFunc("/", homePage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Edgeone API Server - https://edgeone.cloud/api")
}
