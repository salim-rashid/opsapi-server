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

func openrestyRestart (w http.ResponseWriter, r *http.Request) {
        err := exec.Command("sudo", "systemctl", "restart", "openresty").Run()
        if err != nil {
                //fmt.Println("Error starting OpenResty:", err)
               fmt.Fprintf(w, `{"Error restrting OpenResty:"}`,err)
        } else {
                //fmt.Println("OpenResty started successfully")
               fmt.Fprintf(w, `{"status"; 200, response: "Restarted OK" , "healthcheck": "active"}`)
        }
        ////fmt.Printf("%s\n", out)
        ////fmt.Fprintf(w, `{"status": 200, "response": "%s\n"}`, out)
        //fmt.Fprintf(w, `{"status": 200, "response": "Starting Openrest installation....."}`)
}
