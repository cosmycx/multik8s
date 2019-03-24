package main

import (
	"log"
	"net/http"
	"time"
)

const waitSeconds = 20

type server struct {
	router *http.ServeMux
}

func main() {

	router := http.NewServeMux()
	s := server{router: router}

	s.router.HandleFunc("/upload", s.upload)

	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}) // ./

	log.Fatal(http.ListenAndServe(":4500", s.router))

} // .main

// Validate ...
func (s *server) upload(w http.ResponseWriter, r *http.Request) {

	//TODO:
	// w.Header().Add("Access-Control-Allow-Origin", "*")

	//  Upload file
	r.ParseMultipartForm(32 << 20)

	updfile, _, err := r.FormFile("updfile")
	if err != nil {
		log.Printf("Error at upload file: %v\n", err)
		w.Write([]byte("Error at upload file"))
		return
	}
	defer updfile.Close()

	// Simulate work on file :: long duration
	log.Println("Start working on file")

	time.Sleep(waitSeconds * time.Second)

	log.Println("Done working on file")

	w.Write([]byte("file processed, response .."))

} // .Upload
