package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type SomeStruct struct {
	Name  string
	Email string
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("string handler invoked")
	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Println("responseWriter is not really a flusher")
		return
	}
	//this header had no effect
	w.Header().Set("Connection", "Keep-Alive")
	//these two headers are needed to get the http chunk incremently
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	for i := 0; i < 20; i++ {
		// w.Write([]byte("Gorilla! \n"))
		fmt.Println(i)
		fmt.Fprintf(w, "Gorilla! %v \n", i)
		flusher.Flush()
		time.Sleep(1 * time.Second)
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("done")
}

func StringHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("string handler invoked")
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

func logHandler(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("logging middleware start")
		next.ServeHTTP(w, r)
		log.Println("logging middleware ends")
	})
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	sf := http.HandlerFunc(StringHandler)
	r.HandleFunc("/", logHandler(sf))
	r.HandleFunc("/jsonstring", JsonStringHandler)
	r.HandleFunc("/struct", JsonStructHandler)
	r.HandleFunc("/map", JsonMapHandler)
	r.HandleFunc("/stream", StreamHandler)

	// Bind to a port and pass our router in
	log.Println("server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
