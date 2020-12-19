package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func (authApi *AuthApi) MountAuthorizeEndpoint() {
	authPath := authApi.buildAuthorizePath()
	log.Printf("mounting '%s' as authorization path", authPath)
	http.HandleFunc(authPath, authApi.authorize)
}

func (authApi *AuthApi) buildAuthorizePath() string {
	return fmt.Sprintf("%s/authorize", authApi.path)
}

func (authApi *AuthApi) authorize(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	log.Println("building parameters for google sso")

	state, err := getQueryParam(r.URL.RawQuery, "state")
	if err != nil {
		handleError("could not parse query", http.StatusBadRequest, w, r)
		return
	}

	uri := buildAuthorizeUri(
		authApi.oauthCredentials.AuthURI,
		authApi.buildAuthorizeParams(state))

	log.Println("redirecting to google sso")
	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func buildAuthorizeUri(authUri string, params url.Values) string {
	return fmt.Sprintf("%s?%s", authUri, params.Encode())
}

func (authApi *AuthApi) buildAuthorizeParams(state string) url.Values {
	params := url.Values{}
	params.Add("scope", "email")
	params.Add("access_type", "offline")
	params.Add("include_granted_scopes", "true")
	params.Add("response_type", "code")
	params.Add("redirect_uri", authApi.buildRedirectUri())
	params.Add("state", state)
	params.Add("client_id", authApi.oauthCredentials.ClientId)

	log.Println(fmt.Sprintf("redirect_uri is %s", authApi.buildRedirectUri()))
	return params
}

func (authApi *AuthApi) buildRedirectUri() string {
	redirectPort := authApi.redirectPort
	if redirectPort == "" {
		redirectPort = authApi.port
	}

	return fmt.Sprintf("%s://%s:%s%s",
		authApi.protocol,
		authApi.host,
		redirectPort,
		authApi.buildRedirectPath())
}
