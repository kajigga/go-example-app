package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
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

	http.HandleFunc("/version/", logIt(versionHandler))
	http.HandleFunc("/", logIt(homeHandler))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil))
}

func version() {
	fmt.Println(VERSION)
}

type VersionJSON struct {
	Version string `json:"version"`
}
type HomeJSON struct {
	Elements  []int  `json:"elements"`
	PageTitle string `json:"page_title"`
}

func logIt(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
func versionHandler(w http.ResponseWriter, r *http.Request) {
	versionJson := &VersionJSON{Version: VERSION}
	//t, _ := template.ParseFiles("edit.html")
	//t.Execute(w, p)
	js, err := json.Marshal(versionJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// elements := make([]int, 3)
	homeJSON := HomeJSON{Elements: make([]int, 3)}
	homeJSON.PageTitle = "Hello, Goworld!"
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, homeJSON); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
