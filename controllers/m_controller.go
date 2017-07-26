//Package controllers has all the operations related to
package controllers

import (
	"fmt"
	"net/http"
)

//GetVolumeStatus retuns status of the requested volume
func GetVolumeStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"alive" : true}`)
}

//WelcomeMessage returns welcome message
func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"Welcome to OpenEBS s-bot"}`)
}
