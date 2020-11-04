package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/meik99/CoffeeToGO/db"
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
		log.Println(err.Error(), err)
	}
	log.Println(values)

	params := url.Values{}
	params.Add("client_id", credentials.Web.ClientId)
	params.Add("client_secret", credentials.Web.ClientSecret)
	params.Add("code", values.Get("code"))
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", RedirectUri)

	response, err := http.PostForm(credentials.Web.TokenURI, params)
	if err != nil {
		log.Println(err.Error(), err)
	}

	var responseTokens tokens
	responseBodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error(), err)
	}

	err = json.Unmarshal(responseBodyData, &responseTokens)
	if err != nil {
		log.Println(err.Error(), err)
	}

	token, _, err := new(jwt.Parser).ParseUnverified(responseTokens.IdToken, jwt.MapClaims{})
	if err != nil {
		log.Println(err.Error(), err)
	}

	id := ""
	if claims, isMapClaims := token.Claims.(jwt.MapClaims); isMapClaims {
		log.Printf("signed in: '%s'\n", claims["email"])
		id = db.saveTokenForAccount(claims["email"].(string), string(responseBodyData))
	}

	returnParams := url.Values{}
	returnParams.Add("code", id)
	returnUrl, err := url.Parse(values.Get("state"))

	if err != nil {
		log.Println(err.Error(), err)
	} else {
		http.Redirect(w, r, fmt.Sprintf("%s?%s", returnUrl.String(), returnParams.Encode()), http.StatusSeeOther)
	}
}

func authorize(w http.ResponseWriter, r *http.Request) {
	log.Println("building parameters for google sso")
	requestQuery, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println(err.Error(), err)
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
