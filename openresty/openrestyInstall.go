package openresty

import (
	"fmt"
	"net/http"
	//"runtime"
	"os/exec"
	//"os"
	//"bufio"
	//"strings"
	"checkos"
	//"encoding/json"
)

func openrestyInstall(w http.ResponseWriter, r *http.Request) {
	os := checkos.DetectOS(w, r)

	if os == "linux" {
		//fmt.Fprintf(w, `{"%v"}`, os)
		linuxDist := checkos.DetectLinuxDist(w, r)

		//fmt.Fprintf(w, `{"%v"}`, linuxDist)
		switch linuxDist {
		case "Ubuntu":
			//output, err := exec.Command("openresty", "-v").CombinedOutput()
			_, err := exec.Command("openresty", "-v").CombinedOutput()
			if err != nil {
				//w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"status": 200, "response": "OpenResty installation started. Call /v1/openresty_ping to find out the status of the installation"}`)
				//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
				_, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					fmt.Printf("%s", err)
					fmt.Fprintf(w, `{"%s"}`, err)
				}

			} else {
				//w.Header().Set("Content-Type", "application/json")
				//w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"status": 200, "response": "Openresty is already installed"}`)
			}
		case "CentOS":
			//fmt.Printf("Install Openresty on %v.", distro)
			fmt.Fprintf(w, `{"%v"}`, linuxDist)
		case "Amazon Linux":
			//fmt.Printf("Install Openresty on %v.", distro)
			fmt.Fprintf(w, `{"%v"}`, linuxDist)
		default:
			fmt.Fprintf(w, `{"The operating system is not supported!"}`)
		}
		return
	}
}
