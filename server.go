package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var players []Player

func main() {
	mxRouter := mux.NewRouter()
	mxRouter.HandleFunc("/players", options).Methods("OPTIONS")
	mxRouter.HandleFunc("/players/{id}", options).Methods("OPTIONS")
	mxRouter.HandleFunc("/players", getPlayers).Methods("GET")
	mxRouter.HandleFunc("/players", addPlayer).Methods("POST")
	mxRouter.HandleFunc("/players", updatePlayer).Methods("PUT")
	mxRouter.HandleFunc("/players/{id}", deletePlayer).Methods("DELETE")
	mxRouter.HandleFunc("/players/{id}", getPlayer).Methods("GET")
	mxRouter.Methods("GET")
	err := http.ListenAndServe("localhost:8000", mxRouter)
	if err != nil {
		fmt.Printf("Error: ", err)
		return
	}
}

func addPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	id := vars["id"]
	lastName := vars["lastName"]
	newPlayer := Player{ID: id, Name: name, LastName: lastName}
	players = append(players, newPlayer)
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	id := vars["id"]
	i := 0
	for _, currentPlayer := range players {
		if currentPlayer.ID != id {
			players[i] = currentPlayer
			i++
		}
	}
	players = players[:i]
}

func getPlayer(w http.ResponseWriter, r *http.Request) {

}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(players)
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {

}

func init() {
	cristiano := Player{ID: "7", Name: "Cristiano Ronaldo", LastName: "Dos Santos Aveiro"}
	messi := Player{ID: "10", Name: "Lionel", LastName: "Messi"}
	players = []Player{cristiano, messi}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func options(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

type Player struct {
	ID       string
	Name     string
	LastName string
}
