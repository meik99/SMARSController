package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (authApi *AuthApi) MountHealthEndpoint() {
	healthPath := authApi.buildHealthPath()
	log.Printf("mounting '%s' as health path", healthPath)
	http.HandleFunc(healthPath, authApi.health)
}

func (authApi *AuthApi) buildHealthPath() string {
	return fmt.Sprintf("%s/health", authApi.path)
}

func (authApi *AuthApi) health(w http.ResponseWriter, _ *http.Request) {
	enableCors(&w)
	err := sendHealthResponse(w)

	if err != nil {
		log.Println(err.Error(), err)
	}
}

func sendHealthResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(newSuccessResponse())
}
