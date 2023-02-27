package openresty

import (
	"net/http"
)

func OpenrestyConf() {
	http.HandleFunc("/v1/openresty_install", openrestyInstall)
	http.HandleFunc("/v1/openresty_uninstall", openrestyUninstall)
	http.HandleFunc("/v1/openresty_ping", openrestyStatus)
	http.HandleFunc("/v1/openresty_version", openrestyVersion)
	http.HandleFunc("/v1/openresty_start", openrestyStart)
	http.HandleFunc("/v1/openresty_stop", openrestyStop)
	http.HandleFunc("/v1/openresty_restart", openrestyRestart)
	http.HandleFunc("/v1/openresty_reload", openrestyReload)
	http.HandleFunc("/v1/openresty_config_fetch", openrestyconfFetch)
	http.HandleFunc("/v1/openresty_config_upload", openrestyconfUpload)
	http.HandleFunc("/v1/openresty_config_archive", openrestyconfArchive)
        http.HandleFunc("/v1/openresty_config_test", openrestyconfTest)
}

