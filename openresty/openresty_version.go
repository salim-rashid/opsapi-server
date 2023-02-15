package openresty

import (
        //"bufio"
        "fmt"
        "net/http"
        //"os"
        //"runtime"
        "strings"
        "os/exec"
        "encoding/json"
)


type Application struct {
        Name    string `json:"name"`
        Version string `json:"version"`
}

func openrestyVersion (w http.ResponseWriter, r *http.Request) {
        output, err := exec.Command("openresty", "-v").CombinedOutput()
        if err != nil {
                w.Header().Set("Content-Type", "application/json")
                //fmt.Println("OpenResty is not installed")
                fmt.Fprintf(w, `{"OpenResty is not installed"}`)
                return
        }

        versionOutput := strings.Split(string(output), " ")
        version := strings.TrimSpace(versionOutput[2])
        custom_versionOutput := strings.Split(string(version), "/")
        custom_version := strings.TrimSpace(custom_versionOutput[1])

        app := Application{
                Name:    "OpenResty",
                Version: custom_version,
        }

        appJSON, err := json.Marshal(app)
        if err != nil {
                //fmt.Println("Error marshaling JSON:", err)
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintf(w, `{"Error marshaling JSON:"}`, err)
                return
        }

        //fmt.Println(string(appJSON))
        //fmt.Printf(string(appJSON))
        //fmt.Fprintf(w, string(appJSON))
        fmt.Fprintf(w, `{"status": 200, "response": %s"}`, string(appJSON))
}
