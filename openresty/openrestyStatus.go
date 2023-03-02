package openresty

import (
        //"bufio"
        "fmt"
        "net/http"
        //"os"
        //"runtime"
        "strings"
        "os/exec"
        //"encoding/json"
)

func openrestyStatus(w http.ResponseWriter, r *http.Request) {
        //out, err := exec.Command("systemctl", "status", "openresty").CombinedOutput()
        out, _ := exec.Command("systemctl", "status", "openresty").CombinedOutput()
        //if err != nil {
        //      fmt.Fprintf(w, `{"status": 500, "response": "Failed to check Openresty status"}`)
        //      return
        //}
        outStr := string(out)
        lines := strings.Split(outStr, "\n")
        for _, line := range lines {
                if strings.Contains(line, "inactive") {
                        w.WriteHeader(http.StatusOK)
                        splitLine := strings.Split(string(line), ":")
                        output := strings.TrimSpace(splitLine[1])
                        splitOutput := strings.Split(string(output), " ")
                        state := strings.TrimSpace(splitOutput[0])
                        //fmt.Fprintf(w, `{"%s"}`, state)
                        //fmt.Fprintf(w, `{"%s"}`, output)
                        if strings.Contains(state, "inactive") {
                               //fmt.Fprintf(w, `{"status": 200, "response": "Openresty is not running", "healthcheck": "stopped"}`)
                               fmt.Fprintf(w, `{"status": 200, "response": "%v", "healthcheck": "stopped"}`, strings.TrimLeft(strings.Join(splitLine, ":",), " "))
                               //fmt.Fprintf(w, `{"%s"}`, strings.TrimLeft(strings.Join(splitLine, ":",), " "))
                        }//else {
                        //      fmt.Fprintf(w, `{"status": 200, "response": "Openresty is not running", "healthcheck": "inactive"}`)
                        //}
                        break
                }

                //if strings.HasPrefix(line, "     Active:") {
                if strings.Contains(line, "active") {
                        w.WriteHeader(http.StatusOK)
                        splitOutput := strings.Split(string(line), ":")
                        output := strings.TrimSpace(splitOutput[1])
                        if strings.Contains(output, "active") {
                                //fmt.Fprintf(w, `{"status": 200, "response": "Openresty is running", "healthcheck": "active"}`)
                                fmt.Fprintf(w, `{"status": 200, "response": "%v", "healthcheck": "active"}`, strings.TrimLeft(strings.Join(splitOutput, ":",), "     Active:"))
                                //fmt.Fprintf(w, `{"%s"}`, strings.TrimLeft(strings.Join(splitOutput, ":",), " "))
                        }
                        break
                }
       }
}
