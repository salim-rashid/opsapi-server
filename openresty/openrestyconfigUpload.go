package openresty 

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
    //"fileupload"
)



func openrestyconfUpload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
	t, _ := template.ParseFiles("./openresty/upload.html")
	t.Execute(w, nil)
    } else if r.Method == "POST" {

      //fmt.Fprintf(w, "File Upload Endpoint Hit")

      // Parse our multipart form, 10 << 20 specifies a maximum
      // upload of 10 MB files.

      r.ParseMultipartForm(10 << 20)

      // FormFile returns the first file for the given key `myFile`
      // it also returns the FileHeader so we can get the Filename,
      // the Header and the size of the file

      //file, handler, err := r.FormFile("myFile")
      file, _, err := r.FormFile("myFile")
      if err != nil {
         fmt.Println("Error Retrieving the File")
         fmt.Println(err)
         return
      }
      defer file.Close()
      //fmt.Printf("Uploaded File: %+v\n", handler.Filename)
      //fmt.Printf("File Size: %+v\n", handler.Size)
      //fmt.Printf("MIME Header: %+v\n", handler.Header)

      // Create a temporary file within our temp-images directory that follows
      // a particular naming pattern

      //tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
      tempFile, err := ioutil.TempFile("uploads", "upload-*.json")
      if err != nil {
         fmt.Println(err)
      }
      defer tempFile.Close()

      // read all of the contents of our uploaded file into a
      // byte array

      fileBytes, err := ioutil.ReadAll(file)
      if err != nil {
         fmt.Println(err)
      }

      // write this byte array to our temporary file

      tempFile.Write(fileBytes)

      // return that we have successfully uploaded our file!

      fmt.Fprintf(w, `{"status"; 200, response: "Successfully Uploaded File"}`)
   }
}

//-->func setupRoutes() {
//-->    http.HandleFunc("/upload", uploadFile)
//-->    http.ListenAndServe(":8080", nil)
//-->}

//func main() {
//    //-->fmt.Println("Hello World")
//    //-->setupRoutes()
//    http.HandleFunc("/upload", uploadFile)
//    //fileupload.HomePageConf()
//    http.ListenAndServe(":8080", nil)
//}
