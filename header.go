package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var players []Player

func init() {
	cristiano := Player{ID: "7", Name: "Cristiano Ronaldo", LastName: "Dos Santos Aveiro"}
	messi := Player{ID: "10", Name: "Lionel", LastName: "Messi"}
	players = []Player{cristiano, messi}
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pTemplate, _ := template.ParseFiles("players.html")
		err := pTemplate.Execute(w, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(players)

	} else if r.Method == "POST" {
		r.ParseForm()
		playerID := r.FormValue("playerID")
		playerName := r.FormValue("playerName")
		foundPlayer, isFound, _ := updatePlayer(playerID)
		if isFound {
			foundPlayer.Name = playerName
			json.NewEncoder(w).Encode(foundPlayer)
		} else {
			fmt.Fprintf(w, "NOT FOUND")
		}
	} else if r.Method == "DELETE" {
		r.ParseForm()
		playerID := r.FormValue("playerID")
		_, fo, i := updatePlayer(playerID)
		if fo {
			ps := RemoveIndex(players, i)
			json.NewEncoder(w).Encode(ps)
		} else {
			fmt.Fprintf(w, "NOT FOUND")
		}
	}

}

func updatePlayer(id string) (Player, bool, int) {
	var foundPlayer Player
	isFound := false
	index := -1
	for i, player := range players {
		if player.ID == id {
			foundPlayer = player
			isFound = true
			index = i
		}
	}
	return foundPlayer, isFound, index
}
func deletePlayer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	playerID := r.FormValue("playerID")
	_, fo, i := updatePlayer(playerID)
	if fo {
		ps := RemoveIndex(players, i)
		json.NewEncoder(w).Encode(ps)
	} else {
		fmt.Fprintf(w, "NOT FOUND")
	}
}

func RemoveIndex(s []Player, index int) []Player {
	return append(s[:index], s[index+1:]...)
}

type Player struct {
	ID       string
	Name     string
	LastName string
}

func main() {
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/add_cookie", addCookie)
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/upload_file", processFile)
	http.HandleFunc("/players", getPlayers)
	http.HandleFunc("/delete_player", deletePlayer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

func addCookie(w http.ResponseWriter, r *http.Request) {
	g1 := http.Cookie{
		Name:  "myCookie",
		Value: "value"}
	g2 := http.Cookie{
		MaxAge: 5000,
		Name:   "myCookie2",
		Value:  "value2"}
	http.SetCookie(w, &g1)
	http.SetCookie(w, &g2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	value, _ := r.Cookie("myCookie")
	fmt.Fprintln(w, value)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		pTemplate, _ := template.ParseFiles("cookie-form.html")
		err := pTemplate.Execute(w, nil)
		if err != nil {
			fmt.Printf("Error: ", err)
			return
		}
	} else {
		r.ParseForm()
		c1 := http.Cookie{
			Name:  "cookieValue",
			Value: r.FormValue("cookieValue")}
		http.SetCookie(w, &c1)
	}
}

func processFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		pTemplate, _ := template.ParseFiles("file.html")
		err := pTemplate.Execute(w, nil)
		if err != nil {
			fmt.Printf("Error: ", err)
			return
		}
	} else {
		r.ParseMultipartForm(32 << 20)
		file, header, err := r.FormFile("subir")
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "Received Form %v", header.Filename)
		newFile, err := os.OpenFile("./files/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer newFile.Close()
		io.Copy(newFile, file)
	}
}
