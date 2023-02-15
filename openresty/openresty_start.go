package openresty

import (
        //"bufio"
        "fmt"
        "net/http"
        //"os"
        //"runtime"
        //"strings"
        "os/exec"
        //"encoding/json"
)

func openrestyStart (w http.ResponseWriter, r *http.Request) {
        err := exec.Command("sudo", "systemctl", "start", "openresty").Run()
        if err != nil {
                //fmt.Println("Error starting OpenResty:", err)
               fmt.Fprintf(w, `{"Error starting OpenResty:"}`,err)
        } else {
                //fmt.Println("OpenResty started successfully")
               fmt.Fprintf(w, `{"status"; 200, response: "Started OK" , "healthcheck": "active"}`)
        }
        ////fmt.Printf("%s\n", out)
        ////fmt.Fprintf(w, `{"status": 200, "response": "%s\n"}`, out)
        //fmt.Fprintf(w, `{"status": 200, "response": "Starting Openrest installation....."}`)
}
