package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
)

var path = flag.String("path", "./GeoLiteCity", "path to Geo Lite City db")
var port = flag.String("port", "80", "port to listen on")

func main() {
	flag.Parse()

	log.Println("building geo db")
	db, err := NewGeoDB(*path)
	if err != nil {
		log.Fatalf("failed to create geo db: %s", err)
	}
	log.Println("done building geo db")

	http.HandleFunc("/healthcheck", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	})
	http.HandleFunc("/ip/", func(rw http.ResponseWriter, r *http.Request) {
		latlon, err := db.Lookup(strings.TrimLeft(r.URL.Path, "/ip/"))
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
		}
		b, err := json.Marshal(latlon)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
		}
		rw.Write(b)
	})
	log.Printf("listening on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
