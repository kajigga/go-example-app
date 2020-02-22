package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const VERSION = "0.10"

type Args struct {
	GetVersion bool
	Message    string
}

func main() {

	a := Args{}
	flag.BoolVar(&a.GetVersion, "version", false, "get the version")
	flag.StringVar(&a.Message, "message", "", "message to share")

	flag.Parse()

	if a.GetVersion {
		version()
	} else {
		fmt.Println(a.Message)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/version/", middleware(VersionHandler))
	mux.HandleFunc("/", middleware(HomeHandler))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), mux))
}

func version() {
	fmt.Println(VERSION)
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return logIt(f)

}

func logIt(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
