package mdrepo

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func noCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		h.ServeHTTP(w, r)
	})
}

func (mdlook *MDLookManager) StartServer(addr string, secondaryPath string) {
	webUIPath := mdlook.GetWebUIFolderPath()
	indexHtmlPath := mdlook.GetIndexHtmlPath()

	primaryFileServer := noCacheHandler(http.FileServer(http.Dir(webUIPath)))
	secondaryFileServer := noCacheHandler(http.FileServer(http.Dir(secondaryPath)))

	mux := http.NewServeMux()

	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		primaryFilePath := filepath.Join(webUIPath, path)
		if _, err := os.Stat(primaryFilePath); err == nil {
			primaryFileServer.ServeHTTP(w, r)
			return
		}

		secondaryFilePath := filepath.Join(secondaryPath, path)
		if _, err := os.Stat(secondaryFilePath); err == nil {
			secondaryFileServer.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, indexHtmlPath)
	})

	mux.Handle("/", rootHandler)

	fmt.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
