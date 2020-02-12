package main

import (
	"encoding/json"
	"net/http"

	"github.com/nasjp/ginclone"
)

func main() {
	r := ginclone.New()

	r.GET("/hello", HelloWorldHandler)
	r.GET("/users", GetUsersHandler)

	r.Run(":6969")
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
	}{
		{"Luke Skywalker", "male"},
		{"Leia Organa", "female"},
		{"Han Solo", "male"},
		{"Chewbacca", "male"},
	})
}
