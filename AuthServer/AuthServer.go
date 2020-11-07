package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	RedirectUri = "http://localhost:8080/apps/coffeetogo/api/v1/sso/google/redirect"

	CredentialsPath = "client_secret_723945016868-kh3kloreb5uldrneieshhpii939uitgf.apps.googleusercontent.com.json"
)

var credentials googleCredentials

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Println("finalizing authorization")
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println("could not parse query")
		log.Println(err.Error(), err)
		handleError("could not parse query", http.StatusBadRequest, w, r)
		return
	}

	params := url.Values{}
	params.Add("client_id", credentials.Web.ClientId)
	params.Add("client_secret", credentials.Web.ClientSecret)
	params.Add("code", values.Get("code"))
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", RedirectUri)

	response, err := http.PostForm(credentials.Web.TokenURI, params)
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
	params.Add("redirect_uri", RedirectUri)
	params.Add("state", requestQuery.Get("state"))
	params.Add("client_id", credentials.Web.ClientId)
	uri := fmt.Sprintf("%s?%s", credentials.Web.AuthURI, params.Encode())

	log.Println("redirecting to google sso")
	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func test(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("success"))
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

func handleRequests() {
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google", authorize)
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google/redirect", redirect)
	http.HandleFunc("/apps/coffeetogo/api/v1/sso/google/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	readCredentials()

	handleRequests()
}

func readCredentials() {
	credentialsJson, err := os.Open(CredentialsPath)
	if err != nil {
		log.Fatalf(err.Error(), err)
	}
	defer func() {
		err := credentialsJson.Close()
		if err != nil {
			log.Fatalf(err.Error(), err)
		}
	}()

	credentialsJsonData, err := ioutil.ReadAll(credentialsJson)
	if err != nil {
		log.Fatalf(err.Error(), err)
	}

	err = json.Unmarshal(credentialsJsonData, &credentials)
	if err != nil {
		log.Fatalf(err.Error(), err)
	}
}
