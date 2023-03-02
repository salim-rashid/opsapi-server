package checkos

import (
	"fmt"
        "net/http"
        "runtime"
	"os"
	"bufio"
	"strings"
)

//func Ostype() {
//        http.HandleFunc("/ostype", DetectOS)
//}

func DetectOS (w http.ResponseWriter, r *http.Request) string{
        // Detect OS
        o := runtime.GOOS
        switch o {
        case "windows":
                //fmt.Printf("Windows")
                w.Header().Set("Content-Type", "application/json")
                //json.NewEncoder(w).Encode(os)
                fmt.Fprintf(w, `{"message": "Great, you are running %v."}`, o)
        case "darwin":
                // fmt.Println("MAC operating system")
                w.Header().Set("Content-Type", "application/json")
                // json.NewEncoder(w).Encode(os)
                fmt.Fprintf(w, `{"message": "Great, you are running %v."}`, o)
        case "linux":
                // fmt.Println("Linux")
                w.Header().Set("Content-Type", "application/json")
                //json.NewEncoder(w).Encode(os)
                //fmt.Fprintf(w, `{"message": "Great, you are running %v."}`, o)
		return o
                // Detect Linux distribution
                //detectLinuxDist(w, r, o)
                // Install Openresty
                // InstallOpenresty()

        default:
                // fmt.Printf("%s.\n", os)
                //fmt.Printf("The operating system is not supported!")
                w.Header().Set("Content-Type", "application/json")
                fmt.Fprintf(w, `{"message": "The operating system is not supported!"}`)
        }

        // Install Openresty
        // Return a JSON for the installation status
        // fmt.Println(os(w, r))
	return o 
}

func DetectLinuxDist(w http.ResponseWriter, r *http.Request) string  {
        file, err := os.Open("/etc/os-release")
        if err != nil {
                panic(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := scanner.Text()
                if strings.HasPrefix(line, "NAME=") {
                        distro := strings.TrimPrefix(line, "NAME=")
                        distro = strings.Trim(distro, "\"")
                        // fmt.Println("Distribution:", distro)
                        // fmt.Println(distro)
                        w.Header().Set("Content-Type", "application/json")
                        //fmt.Fprintf(w, `{"message": "You are running %v Linux!"}`, distro)
                        //detectLinuxFlavour(w, r, distro)
                        //break
			return distro
                }
        }

	return "" 
}

func DetectLinuxFlavour (w http.ResponseWriter, r *http.Request, distro string) {
        switch distro {
        case "Ubuntu":
                //fmt.Fprintf(w, `{"Install Openresty on %v Linux......\n"}`, distro)
                //ubuntuopenrestyInstall(w, r)
                //installStatus()
                //http.HandleFunc("/v1/openresty_install_status", installStatus())
                //out, err := exec.Command("bash", "-c", "./install_openresty.sh").Output()
                //if err != nil {
                //   fmt.Printf("%s", err)
                //}
                fmt.Fprintf(w, `{"%v"}`, distro)
        case "CentOS":
                //fmt.Printf("Install Openresty on %v.", distro)
                fmt.Fprintf(w, `{"%v"}`, distro)
        case "Amazon Linux":
                //fmt.Printf("Install Openresty on %v.", distro)
                fmt.Fprintf(w, `{"%v"}`, distro)
        default:
                fmt.Fprintf(w, `{"The operating system is not supported!"}`)
	}
	return
 }
