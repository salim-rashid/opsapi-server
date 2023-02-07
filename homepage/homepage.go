package homepage

import (
	"net/http"
)

// func HomePageConf() {
// 	http.HandleFunc("/", homePage)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the Edgeone API Server - https://edgeone.cloud/api")
// }

func HomePageConf() {
	http.Handle("/", http.FileServer(http.Dir("./homepage")))

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err)
	// }
}
