package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/meik99/CoffeeToGO/AuthServer/auth_json"
	"github.com/meik99/CoffeeToGO/AuthServer/credentials"
	"github.com/meik99/CoffeeToGO/AuthServer/db"
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
	enableCors(&w)
	log.Println("finalizing authorization")

	params, err := authApi.buildTokenRequestParams(r.URL.RawQuery)
	if err != nil {
		handleError("could not parse query", http.StatusBadRequest, w, r)
		return
	}

	oauthToken, idToken, err := getCredentialTokens(authApi.oauthCredentials.TokenURI, params)
	if err != nil {
		handleError("could not request auth tokens", http.StatusBadRequest, w, r)
		return
	}

	id := db.SaveTokenForAccount(getEmailFromToken(idToken), *oauthToken)
	redirectToStateFromQueryParams(w, r, id)
}

func redirectToStateFromQueryParams(w http.ResponseWriter, r *http.Request, id string) {
	returnParams := url.Values{}
	returnParams.Add("code", id)

	returnUrl, err := getQueryParam(r.URL.RawQuery, "state")

	if err != nil {
		log.Println("could not parse redirect url")
		log.Println(err.Error(), err)
		handleError("could not parse redirect url", http.StatusInternalServerError, w, r)
		return
	} else {
		http.Redirect(w, r, fmt.Sprintf("%s?%s", returnUrl, returnParams.Encode()), http.StatusSeeOther)
	}
}

func getEmailFromToken(token *jwt.Token) string {
	if claims, isMapClaims := token.Claims.(jwt.MapClaims); isMapClaims {
		return claims["email"].(string)
	}
	return ""
}

func getCredentialTokens(tokenUri string, params url.Values) (*credentials.AuthToken, *jwt.Token, error) {
	response, err := http.PostForm(tokenUri, params)
	if err != nil {
		log.Println("could not request auth tokens")
		log.Println(err.Error(), err)
		return nil, nil, err
	}

	var responseTokens credentials.AuthToken
	err = auth_json.ParseJSONToInterface(response.Body, &responseTokens)
	if err != nil {
		log.Println("could not parse auth token")
		log.Println(err.Error(), err)
		return nil, nil, err
	}

	token, _, err := new(jwt.Parser).ParseUnverified(responseTokens.IdToken, jwt.MapClaims{})
	if err != nil {
		log.Println("could not parse auth token")
		log.Println(err.Error(), err)
		return nil, nil, err
	}

	return &responseTokens, token, nil
}

func (authApi *AuthApi) buildTokenRequestParams(query string) (url.Values, error) {
	code, err := getQueryParam(query, "code")
	if err != nil {
		return url.Values{}, err
	}

	return authApi.buildParamsWithCode(code), nil
}

func (authApi *AuthApi) buildParamsWithCode(code string) url.Values {
	params := url.Values{}
	params.Add("client_id", authApi.oauthCredentials.ClientId)
	params.Add("client_secret", authApi.oauthCredentials.ClientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	params.Add("redirect_uri", authApi.buildRedirectUri())

	return params
}

func getQueryParam(query string, key string) (string, error) {
	values, err := url.ParseQuery(query)
	if err != nil {
		log.Println("could not parse query")
		log.Println(err.Error(), err)
		return "", err
	}
	return values.Get(key), nil
}
