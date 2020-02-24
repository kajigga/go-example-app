package main

import (
	"encoding/json"
	"net/http"
	"path"
	"text/template"
)

func VersionHandler(w http.ResponseWriter, r *http.Request) {
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

func TodosHandler(w http.ResponseWriter, r *http.Request) {

	js, err := json.Marshal(AllTodos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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
