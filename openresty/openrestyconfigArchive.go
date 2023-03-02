package openresty 

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"path/filepath"
	"net/http"
)

//func moveJSONFile(w http.ResponseWriter, r *http.Request) error {
func openrestyconfArchive (w http.ResponseWriter, r *http.Request) {
	// Define the paths of the source and destination files
	//src := "configs/example.json"
	src := "./openresty/config.json"
	dst := "/tmp/config.json"

	// Read the contents of the source file
	bytes, err := ioutil.ReadFile(src)
	if err != nil {
              // Return the error as JSON
              w.Header().Set("Content-Type", "application/json")
              w.WriteHeader(http.StatusInternalServerError)
              json.NewEncoder(w).Encode(map[string]interface{}{
              "error": err.Error(),
              })
	      return 
	      //return err
	}

	// Create the destination directory if it doesn't already exist
	dir := filepath.Dir(dst)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
              w.Header().Set("Content-Type", "application/json")
              w.WriteHeader(http.StatusInternalServerError)
              json.NewEncoder(w).Encode(map[string]interface{}{
              "error": err.Error(),
              })
	      return 
	      //return err
	}

	// Write the contents of the source file to the destination file
	if err = ioutil.WriteFile(dst, bytes, os.ModePerm); err != nil {
	      w.Header().Set("Content-Type", "application/json")
              w.WriteHeader(http.StatusInternalServerError)
              json.NewEncoder(w).Encode(map[string]interface{}{
              "error": err.Error(),
              })
	      return 
	      //return err
	}

	// Send the status of file archiving to the browser.
	w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
        "status":   200,
        "response": "Successfully archived the file " + src,
        })

	// Remove the source file
	if err = os.Remove(src); err != nil {
              w.Header().Set("Content-Type", "application/json")
              w.WriteHeader(http.StatusInternalServerError)
              json.NewEncoder(w).Encode(map[string]interface{}{
              "error": err.Error(),
              })
	      return 
	      //return err
	}

	//return nil
}
