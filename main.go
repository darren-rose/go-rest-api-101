package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var players []Player

func getPlayers(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.Method)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.Method)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range players {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.Method)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range players {
		log.Printf("index %d", index)
		if item.ID == params["id"] {
			if len(players) > 1 {
				players = append(players[:index], players[index+1])
			} else {
				players = append(players[:index])
			}
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.Header.Get("Content-Type"))

	w.Header().Set("Content-Type", "application/json")

	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	player.ID = uuid.New().String()

	players = append(players, player)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&player)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers).Methods(http.MethodGet)
	r.HandleFunc("/players/{id}", getPlayer).Methods(http.MethodGet)
	r.HandleFunc("/players/{id}", deletePlayer).Methods(http.MethodDelete)
	r.HandleFunc("/players", createPlayer).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8000", r))
}
