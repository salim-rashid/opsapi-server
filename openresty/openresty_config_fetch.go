package openresty 

import (
    "fmt"
    "os/exec"
    "io/ioutil"
    "bufio"
    "os"
    "strings"
    "encoding/json"
    "net/http"
)

func saveOpenrestyConfigToFile(w http.ResponseWriter, r *http.Request, filename string) error {

    // Use sudo to run the command as root
    cmd := exec.Command("sudo", "openresty", "-T")
    output, err := cmd.Output()
    if err != nil {
        return err
	//fmt.Fprintf(w, `{"Error restrting OpenResty:"}`,err)
    }

    // Write the output to the specified file
    err = ioutil.WriteFile(filename, output, 0644)
    if err != nil {
        return err
    }

    //fmt.Printf("Openresty configuration saved to %s\n", filename)
    convertFileOutputtoJson(w, r, filename)

    return nil
}

func convertFileOutputtoJson(w http.ResponseWriter, r *http.Request, filename string) {
    file, err := os.Open("output")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    var result map[string]interface{}
    result = make(map[string]interface{})

    scanner := bufio.NewScanner(file)
    section := ""

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        // Identify sections
        if strings.HasSuffix(line, "{") {
            section = strings.TrimSpace(strings.TrimSuffix(line, "{"))
            result[section] = make(map[string]interface{})
            continue
        } else if line == "}" {
            section = ""
            continue
        }

        // Extract values from lines
        parts := strings.SplitN(line, " ", 2)
        key := strings.TrimSpace(parts[0])
        //value := strings.TrimSpace(parts[1])
        if len(parts) > 1 {
            value := strings.TrimSpace(parts[1])
            // Add values to the appropriate section
            if section == "" {
               result[key] = value
            } else {
                sectionMap := result[section].(map[string]interface{})
                sectionMap[key] = value
            }
        }
    }


    // Convert the result to JSON and print it
    jsonResult, err := json.MarshalIndent(result, "", "  ")
    if err != nil {
        fmt.Println(err)
        return
    }

    //fmt.Println(string(jsonResult))
    //fmt.Fprintln(w, `{%v}`, string(jsonResult))
    fmt.Fprintln(w, string(jsonResult))
}

func openrestyconfFetch(w http.ResponseWriter, r *http.Request) {
	filename := "./output"
	saveOpenrestyConfigToFile(w, r, filename)

}
