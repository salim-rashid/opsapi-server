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

func openrestyUninstall(w http.ResponseWriter, r *http.Request) {
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
				//fmt.Fprintf(w, `{"status": 200, "response": "OpenResty uninstallation started\n"}`)
				fmt.Fprintf(w, `{"Openrety is not installed!: %v\n"}`, err)
				return
				//fmt.Println("OpenResty is not installed")
				//return
			} else {
				//w.Header().Set("Content-Type", "application/json")
				//w.WriteHeader(http.StatusOK)
				//fmt.Printf("Openresty is already installed.\n")
				//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
				_, err := exec.Command("bash", "-c", "./openresty/uninstall_openresty.sh").Output()
				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, `{"Openresty is not installed. %s"}`, err)

				}
				fmt.Fprintf(w, `{"status": 200, "response": "Successfully uninstalled Openresty"}`)
				//installStatus(w, r)
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
