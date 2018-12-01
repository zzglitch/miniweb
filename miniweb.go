package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type fileHandlerWithCacheControl struct {
	fileServer http.Handler
}

func (f *fileHandlerWithCacheControl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	f.fileServer.ServeHTTP(w, r)
}

func main() {
	dir := flag.String("dir", "./www", "Directory containing the files to serve")
	port := flag.Int("port", 8080, "Port to run the web server")
	flag.Parse()

	http.Handle("/", &fileHandlerWithCacheControl{fileServer: http.FileServer(http.Dir(*dir))})

	log.Println("Starting web server on port " + strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
