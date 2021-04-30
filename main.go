package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	ID      string
	Name    string
	Age     int
	Address string
	Dom     string
}

var datas = []Data{
	{"D001", "afis", 23, "lumajang", "lumajang"},
	{"D002", "afis2", 23, "lumajang", "lumajang"},
	{"D003", "afis3", 23, "lumajang", "lumajang"},
	{"D004", "afis4", 23, "lumajang", "lumajang"},
	{"D005", "afis5", 23, "lumajang", "lumajang"},
}

func main() {
	// membuat webservice dulu
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			byteJson, err := json.Marshal(datas)

			if err != nil {
				http.Error(w, "error internal server", http.StatusInternalServerError)
				return
			}

			w.Write(byteJson)
			return
		}

		http.Error(w, "error not method POST", http.StatusInternalServerError)
	})

	port := "localhost:4444"

	fmt.Println("server running on port", port)

	http.ListenAndServe(port, nil)
}
