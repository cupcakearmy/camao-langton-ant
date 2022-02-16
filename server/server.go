package server

import (
	"encoding/json"
	"io/ioutil"
	"langton/logic"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func preflight(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Init() {
	port := "8000"

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			preflight(&w)
			return
		}
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		enableCors(&w)
		world := logic.RandomWorld(10)
		json.NewEncoder(w).Encode(world)
	})

	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			preflight(&w)
			return
		}
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		enableCors(&w)
		reqBody, _ := ioutil.ReadAll(r.Body)
		var world logic.World
		json.Unmarshal(reqBody, &world)
		logic.NextWorldState(&world)
		json.NewEncoder(w).Encode(world)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
