package main

import (
	"fmt"
	"net/http"

	"ascii-art-web/source"
)

const PORT = ":8080"

func main() {
	fmt.Println("http://localhost" + PORT)
	http.HandleFunc("/", source.MainHandler)
	http.HandleFunc("/ascii-art-web", source.HandleAscii)
	err := http.ListenAndServe(PORT, nil)
	source.CheckError(err)
}
