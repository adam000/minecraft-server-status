package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/adam000/goutils/page"
	"github.com/adam000/goutils/shell"
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

type config struct {
	AddressAndPort string
	Password       string
}

type status struct {
	ErrorText        string
	IsOnline         bool
	NumPlayersOnline int
	PlayersOnline    []string
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			json.NewEncoder(w).Encode(status{
				ErrorText: fmt.Sprintf("Error getting status: %s", rec.(error)),
			})
		}
	}()

	var cfg config
	configFile, err := os.Open("config.json")
	if err != nil {
		json.NewEncoder(w).Encode(status{
			ErrorText: fmt.Sprintf("Error loading config file: %s", err),
		})
		return
	}
	parser := json.NewDecoder(configFile)
	parser.Decode(cfg)

	stdout, stderr, err := shell.RunInDir(".", "/bin/Batchcraft", "-a", cfg.AddressAndPort, "-p", cfg.Password, "-c", "list")

	if err != nil {
		log.Printf("Error running batchcraft: %s (%s)", err, stderr)
		json.NewEncoder(w).Encode(status{
			ErrorText: fmt.Sprintf("Could not connect: %s", err),
		})
		return
	}

	if stderr != "" {
		log.Printf("Error text written while running batchcraft: %s", stderr)
		json.NewEncoder(w).Encode(status{
			ErrorText: fmt.Sprintf("Could not connect: %s", stderr),
		})
		return
	}

	if stdout == "" {
		log.Printf("Got no output from command!")
		json.NewEncoder(w).Encode(status{
			ErrorText: "Got no response",
		})
		return
	}

	// The first line of output says "There are x out of a maximum y players online.\n<list of players>". Get the number out and return that
	re := regexp.MustCompile("(§.)")
	stdout = re.ReplaceAllString(stdout, "")

	lines := strings.Split(stdout, "\n")
	statusLine := strings.Split(lines[0], " ")
	numPlayers, err := strconv.Atoi(statusLine[2])
	if err != nil {
		json.NewEncoder(w).Encode(status{
			ErrorText: fmt.Sprintf("Internal error splitting strings: %s", err),
		})
		return
	}

	if numPlayers == 0 {
		json.NewEncoder(w).Encode(status{
			IsOnline:         true,
			NumPlayersOnline: numPlayers,
		})
		return
	}

	players := make([]string, len(lines)-2)
	for i, player := range lines[1 : len(lines)-1] {
		players[i] = strings.SplitAfter(player, " ")[1]
	}

	json.NewEncoder(w).Encode(status{
		IsOnline:         true,
		NumPlayersOnline: numPlayers,
		PlayersOnline:    players,
	})
}
