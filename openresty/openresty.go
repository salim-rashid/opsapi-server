package openresty

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"os/exec"
	"encoding/json"
)

func OpenrestyConf() {
	http.HandleFunc("/v1/openresty_install", openrestyInstall)
	http.HandleFunc("/v1/openresty_ping", openrestyStatus)
	http.HandleFunc("/v1/openresty_version", openrestyVersion)
	http.HandleFunc("/v1/openresty_start", openrestyStart)
	http.HandleFunc("/v1/openresty_stop", openrestyStop)
	http.HandleFunc("/v1/openresty_restart", openrestyRestart)
	http.HandleFunc("/v1/openresty_reload", openrestyReload)

	// http.HandleFunc("/v1/openresty_uninstall", OpenrestyUninstall)

	// http.HandleFunc("/v1/openresty_update", OpenrestyUpdate)

	// http.HandleFunc("/v1/openresty_start", OpenrestyStart)
}

// type Openresty struct {
// 	Name    string `json:"name"`
// 	Version int    `json:"version"`
// }

func openrestyInstall(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Do Openresty installation.")

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
		// Detect Linux distribution
		DetectLinuxDist(w, r, o)
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
}

func DetectLinuxDist(w http.ResponseWriter, r *http.Request, o string) {
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
			InstallOpenresty(w, r, distro)
			break
		}
	}
}

 func InstallOpenresty (w http.ResponseWriter, r *http.Request, distro string) {
 	switch distro {
 	case "Ubuntu":
 		//fmt.Fprintf(w, `{"Install Openresty on %v Linux......\n"}`, distro)
 		deployOpenresty(w, r)
 		//installStatus()
 		//http.HandleFunc("/v1/openresty_install_status", installStatus())
 		//out, err := exec.Command("bash", "-c", "./install_openresty.sh").Output()
 		//if err != nil {
 		//   fmt.Printf("%s", err)
 		//}
 		//fmt.Printf("%s", out)
 	case "CentOS":
 		fmt.Printf("Install Openresty on %v.", distro)
 	case "Amazon Linux":
 		fmt.Printf("Install Openresty on %v.", distro)
 	default:
 		fmt.Println("The operating system is not supported!.")
 	}
 }

 func deployOpenresty (w http.ResponseWriter, r *http.Request) {

 	//output, err := exec.Command("openresty", "-v").CombinedOutput()
 	_, err := exec.Command("openresty", "-v").CombinedOutput()
 	if err != nil {
		//w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": 200, "response": "OpenResty installation started. Call /v1/openresty_ping to find out the status of the installation"}`)
 		//fmt.Println("OpenResty is not installed")
 		//return
 		//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
 		_, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
 		if err != nil {
		        w.Header().Set("Content-Type", "application/json")
 			fmt.Printf("%s", err)
			fmt.Fprintf(w, `{"%s"}`, err)
 		
		}

		//installStatus(w, r)

 		//fmt.Printf("%s", out)
 	} else {
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusOK)
 	        //fmt.Printf("Openresty is already installed.\n")
		fmt.Fprintf(w, `{"status": 200, "response": "Openresty is already installed"}`)
		//installStatus(w, r)
 	}

 	//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()

 	//if err != nil {
 	//   fmt.Printf("%s", err)
 	//}

 	//fmt.Printf("%s", out)
 	////return

 }

 type Application struct {
 	Name    string `json:"name"`
 	Version string `json:"version"`
 }

 // func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
 //      fmt.Fprintf(w, "Do Openresty uninstallation.")
 // }

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
                               fmt.Fprintf(w, `{"status": 200, "response": %v, "healthcheck": "stopped"}`, output)
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
                                fmt.Fprintf(w, `{"status": 200, "response": %v, "healthcheck": "active"}`, output)
                        }
                        break
                }
       }
}

func openrestyStart (w http.ResponseWriter, r *http.Request) {
        err := exec.Command("sudo", "systemctl", "start", "openresty").Run()
	if err != nil {
                //fmt.Println("Error starting OpenResty:", err)
               fmt.Fprintf(w, `{"Error starting OpenResty:"}`,err)
        } else {
                //fmt.Println("OpenResty started successfully")
               fmt.Fprintf(w, `{"status"; 200, response: "Started OK" , "healthcheck": "active"}`)
        }
        ////fmt.Printf("%s\n", out)
        ////fmt.Fprintf(w, `{"status": 200, "response": "%s\n"}`, out)
        //fmt.Fprintf(w, `{"status": 200, "response": "Starting Openrest installation....."}`)
}

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

//func main() {
//        err := exec.Command("sudo", "systemctl", "start", "openresty").Run()
//        //if err != nil {
//        //      return err
//        //}
//
//        //return nil
//        if err != nil {
//                fmt.Println("Error starting OpenResty:", err)
//        } else {
//                fmt.Println("OpenResty started successfully")
//        }



//}

// // func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty uninstallation.")
// // }

// // func OpenrestyUpdate(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty update.")
// // }

// // func OpenrestyStart(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty restart.")
// // }
