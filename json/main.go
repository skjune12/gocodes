package main

import (
	"encoding/json"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}

	response := Response{
		Status: "ok",
		Data:   "success",
	}

	json, _ := json.Marshal(response)

	w.Header(200)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
