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




func openrestyReload (w http.ResponseWriter, r *http.Request) {
        err := exec.Command("sudo", "systemctl", "reload", "openresty").Run()
        if err != nil {
                //fmt.Println("Error starting OpenResty:", err)
               fmt.Fprintf(w, `{"Error reloading OpenResty:"}`,err)
        } else {
                //fmt.Println("OpenResty started successfully")
               fmt.Fprintf(w, `{"status"; 200, response: "Reloaded OK" , "healthcheck": "active"}`)
        }
        ////fmt.Printf("%s\n", out)
        ////fmt.Fprintf(w, `{"status": 200, "response": "%s\n"}`, out)
        //fmt.Fprintf(w, `{"status": 200, "response": "Starting Openrest installation....."}`)
}
