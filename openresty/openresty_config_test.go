package openresty

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "net/http"
)

type OpenrestyResult struct {
    Status 	   string `json:"status"`
    SyntaxStatus   string `json:"syntax_status"`
    TestStatus     string `json:"test_status"`
    //ConfigFilePath string `json:"config_file_path"`
}

func openrestyconfTest(w http.ResponseWriter, r *http.Request) {

    }

    out, err := exec.Command("sudo", "openresty", "-t").CombinedOutput()
    if err != nil {
        fmt.Fprintf(w, `{"Error running command: "}`, err)
    }
    result := OpenrestyResult{}
    if string(out) != "" {
	result.Status = "200"
        result.SyntaxStatus = "ok"
        result.TestStatus = "successful"
        //result.ConfigFilePath = "/usr/local/openresty/nginx/conf/nginx.conf"
    } else {
	result.Status = "500"
        result.SyntaxStatus = "error"
        result.TestStatus = "failed"
        //result.ConfigFilePath = ""
    }
    jsonResult, err := json.Marshal(result)
    if err != nil {
        fmt.Fprintf(w, `{"Error marshalling JSON: "}`, err)
    }
    fmt.Fprintf(w, `%v`, (string(jsonResult)))
}

