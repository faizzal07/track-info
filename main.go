package main

import (
	"net/http"
	"tracks/handler"
)

func main() {
	http.HandleFunc("/getInfo", handler.GetInfoHandler)
	http.ListenAndServe(":8080", nil)
}
