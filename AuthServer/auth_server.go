package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/meik99/CoffeeToGO/AuthServer/credentials"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	CredentialsPath = "client_secret_723945016868-kh3kloreb5uldrneieshhpii939uitgf.apps.googleusercontent.com.json"
)

var oauthCredentials credentials.OAuthCredentials

var host = os.Getenv("COFFEE_AUTH_HOST")
var port = os.Getenv("COFFEE_AUTH_PORT")
var protocol = os.Getenv("COFFEE_AUTH_PROTOCOL")
var path = os.Getenv("COFFEE_AUTH_PATH")
var mainUri = fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
var redirectPath = fmt.Sprintf("%s/redirect", path)
var healthPath = fmt.Sprintf("%s/health", path)
var redirectUri = fmt.Sprintf("%s/redirect", mainUri)

func redirect(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("finalizing authorization")
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("could not parse query")
		log.Println(err.Error(), err)
		handleError("could not parse query", http.StatusBadRequest, w, r)
		return
	}

	params := url.Values{}
	params.Add("client_id", oauthCredentials.ClientId)
	params.Add("client_secret", oauthCredentials.ClientSecret)
	params.Add("code", values.Get("code"))
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", redirectUri)

	response, err := http.PostForm(oauthCredentials.TokenURI, params)
	if err != nil {
		log.Println("could not request auth tokens")
		log.Println(err.Error(), err)
		handleError("could not request auth tokens", http.StatusBadRequest, w, r)
		return
	}

	var responseTokens tokens
	responseBodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("could not read authorization request")
		log.Println(err.Error(), err)
		handleError("could not read authorization request", http.StatusInternalServerError, w, r)
		return
	}

	err = json.Unmarshal(responseBodyData, &responseTokens)
	if err != nil {
		log.Println("could not parse auth token")
		log.Println(err.Error(), err)
		handleError("could not parse auth token", http.StatusInternalServerError, w, r)
		return
	}

	token, _, err := new(jwt.Parser).ParseUnverified(responseTokens.IdToken, jwt.MapClaims{})
	if err != nil {
		log.Println("could not parse auth token")
		log.Println(err.Error(), err)
		handleError("could not parse auth token", http.StatusInternalServerError, w, r)
		return
	}

	id := ""
	if claims, isMapClaims := token.Claims.(jwt.MapClaims); isMapClaims {
		id = saveTokenForAccount(claims["email"].(string), responseTokens)
	}

	returnParams := url.Values{}
	returnParams.Add("code", id)
	returnUrl, err := url.Parse(values.Get("state"))

	if err != nil {
		log.Println("could not parse redirect url")
		log.Println(err.Error(), err)
		handleError("could not parse redirect url", http.StatusInternalServerError, w, r)
		return
	} else {
		http.Redirect(w, r, fmt.Sprintf("%s?%s", returnUrl.String(), returnParams.Encode()), http.StatusSeeOther)
	}
}

func authorize(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("building parameters for google sso")
	requestQuery, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("could not parse query")
		log.Println(err.Error(), err)
		handleError("could not parse query", http.StatusBadRequest, w, r)
		return
	}

	params := url.Values{}
	params.Add("scope", "email")
	params.Add("access_type", "offline")
	params.Add("include_granted_scopes", "true")
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirectUri)
	params.Add("state", requestQuery.Get("state"))
	params.Add("client_id", oauthCredentials.ClientId)
	uri := fmt.Sprintf("%s?%s", oauthCredentials.AuthURI, params.Encode())

	log.Println("redirecting to google sso")
	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func health(w http.ResponseWriter, _ *http.Request) {
	enableCors(&w)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
	err := json.NewEncoder(w).Encode(response{
		Status:  200,
		Message: "success",
	})
	if err != nil {
		log.Println(err.Error(), err)
	}
}

func handleError(msg string, errorCode int, w http.ResponseWriter, r *http.Request) {
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

func handleRequests() {
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

	log.Printf("using '%s' as host", host)
	log.Printf("using '%s' as port", port)
	log.Printf("using '%s' as protocol", protocol)
	log.Printf("using '%s' as main path", path)

	log.Printf("mounting '%s' as authorization path", path)
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google", authorize)

	log.Printf("mounting '%s' as redirect path", redirectPath)
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google/redirect", redirect)

	log.Printf("mounting '%s' as health path", healthPath)
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google/health", health)

	log.Printf("mounting app on port '%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	readCredentials()

	handleRequests()
}

func readCredentials() {
	credentialsProvider := credentials.NewGoogleFileProvider(CredentialsPath)
	oauth, err := credentialsProvider.GetCredentials()

	if err != nil {
		log.Fatalln(err)
	}

	oauthCredentials = oauth
}
