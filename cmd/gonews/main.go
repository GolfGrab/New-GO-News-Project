package main

import (
	"net/http"

	"github.com/GolfGrab/New-GO-News-Project/pkg/app"
)

const port = ":8080"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
