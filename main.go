// main.go

package main

import (
	"encoding/json"
	_ "explore/docs"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Simple API
// @version 1.0
// @description This is a simple API.
// @host localhost:8080
// @BasePath /

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// greet returns a simple greeting message
// @Summary Greet
// @Description Respond with a greeting message
// @Tags greet
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/greet [get]
func greet(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	err := enc.Encode(map[string]string{"message": "Hello, World!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
