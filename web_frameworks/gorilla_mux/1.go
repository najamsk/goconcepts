package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type SomeStruct struct {
	Name  string
	Email string
}

func StringHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func JsonStringHandler(w http.ResponseWriter, r *http.Request) {
	j := `{"name": "najam awan", "email":"najamsk@gmail.com"}`
	w.Write([]byte(j))
}

func JsonStructHandler(w http.ResponseWriter, r *http.Request) {
	data := SomeStruct{Name: "najam", Email: "najamsk@gmail.com"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func JsonMapHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	resp["topic"] = "user/request"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", StringHandler)
	r.HandleFunc("/jsonstring", JsonStringHandler)
	r.HandleFunc("/struct", JsonStructHandler)
	r.HandleFunc("/map", JsonMapHandler)

	// Bind to a port and pass our router in
	log.Println("server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
