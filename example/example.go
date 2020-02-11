package main

import (
	"encoding/json"
	"net/http"

	"github.com/nasjp/ginclone"
)

func main() {
	r := ginclone.New()

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	})

	r.GET("/users", func(w http.ResponseWriter, req *http.Request) {
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
	})

	r.Run(":6969")
}
