package api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func (authApi *AuthApi) MountAuthorizeEndpoint() {
	log.Printf("mounting '%s' as authorization path", authApi.path)
	http.HandleFunc(authApi.path, authApi.authorize)
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

	return params
}

func (authApi *AuthApi) buildRedirectUri() string {
	return fmt.Sprintf("%s://%s:%s%s",
		authApi.protocol,
		authApi.host,
		authApi.port,
		authApi.path)
}
