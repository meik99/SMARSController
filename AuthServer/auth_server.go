package main

import (
	"github.com/meik99/CoffeeToGO/AuthServer/api"
	"github.com/meik99/CoffeeToGO/AuthServer/credentials"
	"log"
	"os"
)

const (
	CredentialsPath = "client_secret_723945016868-kh3kloreb5uldrneieshhpii939uitgf.apps.googleusercontent.com.json"
)

var host = os.Getenv("COFFEE_AUTH_HOST")
var port = os.Getenv("COFFEE_AUTH_PORT")
var protocol = os.Getenv("COFFEE_AUTH_PROTOCOL")
var path = os.Getenv("COFFEE_AUTH_PATH")
var redirectPort = os.Getenv("COFFEE_AUTH_REDIRECT_PORT")

func handleRequests() {
	exitIfEnvVarsAreMissing()
	logEnvVarValues()

	authApi := api.NewAuthApi(protocol, host, port, path, loadCredentials(), redirectPort)
	authApi.MountAuthorizeEndpoint()
	authApi.MountRedirectEndpoint()
	authApi.MountHealthEndpoint()
	authApi.Start()
}

func main() {
	handleRequests()
}

func logEnvVarValues() {
	log.Printf("using '%s' as host", host)
	log.Printf("using '%s' as port", port)
	log.Printf("using '%s' as protocol", protocol)
	log.Printf("using '%s' as main path", path)
}

func exitIfEnvVarsAreMissing() {
	if host == "" {
		log.Fatalln("environment variable COFFEE_AUTH_HOST not set")
	}
	if port == "" {
		log.Fatalln("environment variable COFFEE_AUTH_PORT not set")
	}
	if protocol == "" {
		log.Fatalln("environment variable COFFEE_AUTH_PROTOCOL not set")
	}
	if path == "" {
		log.Fatalln("environment variable COFFEE_AUTH_PATH not set")
	}
}

func loadCredentials() credentials.OAuthCredentials {
	envCredentials := os.Getenv("COFFEE_AUTH_CREDENTIALS")
	credentialsProvider := createSourceDependentCredentialsProvider(envCredentials)
	oauth, err := credentialsProvider.GetCredentials()

	if err != nil {
		log.Fatalln(err)
	}

	return oauth
}

func createSourceDependentCredentialsProvider(envCredentials string) credentials.Provider {
	if envCredentials != "" {
		log.Println("using credentials from environment")
		return credentials.NewGoogleEnvironmentProvider("COFFEE_AUTH_CREDENTIALS")
	} else {
		log.Println("using credentials from file")
		return credentials.NewGoogleFileProvider(CredentialsPath)
	}
}
