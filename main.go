package main

import (
	"fmt"
	"net/http"

	"github.com/satyamz/s-bot/controllers"
)

func main() {
	http.HandleFunc("/status", controllers.GetVolumeStatus)
	http.HandleFunc("/", controllers.WelcomeMessage)
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
