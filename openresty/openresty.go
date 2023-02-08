package openresty

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func OpenrestyConf() {
	http.HandleFunc("/v1/openresty_install", OpenrestyInstall)

	// http.HandleFunc("/v1/openresty_uninstall", OpenrestyUninstall)

	// http.HandleFunc("/v1/openresty_update", OpenrestyUpdate)

	// http.HandleFunc("/v1/openresty_start", OpenrestyStart)
}

// type Openresty struct {
// 	Name    string `json:"name"`
// 	Version int    `json:"version"`
// }

func OpenrestyInstall(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, `{"message": "Great, you are running %v."}`, o)
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
			fmt.Fprintf(w, `{"message": "You are running %v Linux!"}`, distro)
			// InstallOpenresty(distro)
			break
		}
	}
}

// func InstallOpenresty(distro string) {
// 	switch distro {
// 	case "Ubuntu":
// 		//fmt.Printf("Install Openresty on %v.\n", distro)
// 		deployOpenresty()
// 		installStatus()
// 		//http.HandleFunc("/v1/openresty_install_status", installStatus())
// 		//out, err := exec.Command("bash", "-c", "./install_openresty.sh").Output()
// 		//if err != nil {
// 		//   fmt.Printf("%s", err)
// 		//}
// 		//fmt.Printf("%s", out)
// 	case "CentOS":
// 		fmt.Printf("Install Openresty on %v.", distro)
// 	case "Amazon Linux":
// 		fmt.Printf("Install Openresty on %v.", distro)
// 	default:
// 		fmt.Println("The operating system is not supported!.")
// 	}
// }

// func deployOpenresty() {

// 	//output, err := exec.Command("openresty", "-v").CombinedOutput()
// 	_, err := exec.Command("openresty", "-v").CombinedOutput()
// 	if err != nil {
// 		fmt.Println("OpenResty is not installed")
// 		fmt.Println("The installation will now begin ..........")
// 		//return
// 		//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
// 		_, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()
// 		if err != nil {
// 			fmt.Printf("%s", err)
// 		}

// 		//fmt.Printf("%s", out)
// 	} //else {
// 	//   fmt.Printf("Openresty is already installed.\n")
// 	//}

// 	//out, err := exec.Command("bash", "-c", "./openresty/install_openresty.sh").Output()

// 	//if err != nil {
// 	//   fmt.Printf("%s", err)
// 	//}

// 	//fmt.Printf("%s", out)
// 	////return

// }

// type Application struct {
// 	Name    string `json:"name"`
// 	Version string `json:"version"`
// }

// // func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
// //      fmt.Fprintf(w, "Do Openresty uninstallation.")
// // }

// func installStatus() {
// 	output, err := exec.Command("openresty", "-v").CombinedOutput()
// 	if err != nil {
// 		fmt.Println("OpenResty is not installed")
// 		return
// 	}

// 	versionOutput := strings.Split(string(output), " ")
// 	version := strings.TrimSpace(versionOutput[2])
// 	custom_versionOutput := strings.Split(string(version), "/")
// 	custom_version := strings.TrimSpace(custom_versionOutput[1])

// 	app := Application{
// 		Name:    "OpenResty",
// 		Version: custom_version,
// 	}

// 	appJSON, err := json.Marshal(app)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	//fmt.Println(string(appJSON))
// 	fmt.Printf(string(appJSON))
// }

// // func OpenrestyUninstall(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty uninstallation.")
// // }

// // func OpenrestyUpdate(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty update.")
// // }

// // func OpenrestyStart(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Do Openresty restart.")
// // }
