package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", "apple")
}

////go:embed hello.txt
//var s string

//go:embed build
var embeddedFiles embed.FS

func getFileSystem() http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	// print(s)

	r := mux.NewRouter()

	// http.Handle("/", http.FileServer(getFileSystem()))

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/build/").Handler(http.StripPrefix("/build/", http.FileServer(getFileSystem())))

	r.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("listening to port :8000")

	log.Fatal(srv.ListenAndServe())
}
