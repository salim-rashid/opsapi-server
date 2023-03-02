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

func openrestyStop (w http.ResponseWriter, r *http.Request) {
        err := exec.Command("sudo", "systemctl", "stop", "openresty").Run()
        if err != nil {
                //fmt.Println("Error starting OpenResty:", err)
               fmt.Fprintf(w, `{"Error stopping OpenResty:"}`,err)
        } else {
                //fmt.Println("OpenResty started successfully")
               fmt.Fprintf(w, `{"status"; 200, response: "Stopped OK" , "healthcheck": "stopped"}`)
        }
        ////fmt.Printf("%s\n", out)
        ////fmt.Fprintf(w, `{"status": 200, "response": "%s\n"}`, out)
        //fmt.Fprintf(w, `{"status": 200, "response": "Starting Openrest installation....."}`)
}
