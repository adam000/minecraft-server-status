package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/adam000/goutils/page"
	"github.com/gorilla/mux"
)

var tpl *template.Template

func main() {
	tpl = template.New("T")
	template.Must(tpl.ParseGlob("templates/*"))

	addHandlers()
}

func addHandlers() {
	page.SetSiteTitle("Minecraft Server Status")
	r := mux.NewRouter()

	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/status", statusHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var page = page.NewPage()
	page.SetTitle("MC Server Status")

	tpl.ExecuteTemplate(w, "page_home.html", page)
}

type status struct {
	ErrorText     string
	IsOnline      bool
	PlayersOnline int
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if rand.Float64() < 0.1 {
		status := status{
			ErrorText:     "Could not connect",
			IsOnline:      false,
			PlayersOnline: 0,
		}

		json.NewEncoder(w).Encode(status)
	} else if rand.Float64() < 0.5 {
		status := status{
			IsOnline:      true,
			PlayersOnline: 2,
		}

		json.NewEncoder(w).Encode(status)
	} else {
		status := status{
			IsOnline: false,
		}

		json.NewEncoder(w).Encode(status)
	}
}
