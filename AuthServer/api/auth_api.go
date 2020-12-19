package api

import (
	"encoding/json"
	"fmt"
	"github.com/meik99/CoffeeToGO/AuthServer/credentials"
	"log"
	"net/http"
)

type Api interface {
	MountAuthorizeEndpoint()
	MountRedirectEndpoint()
	MountHealthEndpoint()
	Start()
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func newSuccessResponse() response {
	return response{
		Status:  200,
		Message: "success",
	}
}

type AuthApi struct {
	protocol         string
	host             string
	port             string
	path             string
	redirectPort     string
	oauthCredentials credentials.OAuthCredentials
}

func NewAuthApi(protocol, host, port, path string, oauthCredentials credentials.OAuthCredentials, redirectPort string) Api {
	return &AuthApi{
		protocol:         protocol,
		host:             host,
		port:             port,
		path:             path,
		redirectPort:     redirectPort,
		oauthCredentials: oauthCredentials,
	}
}

func (authApi *AuthApi) Start() {
	log.Printf("mounting app on port '%s", authApi.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", authApi.port), nil))
}

func handleError(msg string, errorCode int, w http.ResponseWriter, _ *http.Request) {
	type errorMessage struct {
		message string
		code    int
	}

	w.WriteHeader(errorCode)
	err := json.NewEncoder(w).Encode(errorMessage{
		message: msg,
		code:    errorCode,
	})
	if err != nil {
		log.Println("could not send error message")
		log.Println(err.Error(), err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("error handling request"))
		if err != nil {
			log.Println(err.Error(), err)
		}
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
