package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (authApi *AuthApi) MountRedirectEndpoint() {
	redirectPath := authApi.buildRedirectPath()
	log.Printf("mounting '%s' as redirect path", redirectPath)
	http.HandleFunc(redirectPath, authApi.redirect)
}

func (authApi *AuthApi) buildRedirectPath() string {
	return fmt.Sprintf("%s/redirect", authApi.path)
}

func (authApi *AuthApi) redirect(w http.ResponseWriter, r *http.Request) {
	redirectPath := authApi.buildRedirectPath()
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
	params.Add("client_id", authApi.oauthCredentials.ClientId)
	params.Add("client_secret", authApi.oauthCredentials.ClientSecret)
	params.Add("code", values.Get("code"))
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", redirectPath)

	response, err := http.PostForm(authApi.oauthCredentials.TokenURI, params)
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
