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
	redirectUri := authApi.buildRedirectUri()

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
	params.Add("client_id", authApi.oauthCredentials.ClientId)
	uri := fmt.Sprintf("%s?%s", authApi.oauthCredentials.AuthURI, params.Encode())

	log.Println("redirecting to google sso")
	http.Redirect(w, r, uri, http.StatusSeeOther)
}

func (authApi *AuthApi) buildRedirectUri() string {
	return fmt.Sprintf("%s://%s:%s%s",
		authApi.protocol,
		authApi.host,
		authApi.port,
		authApi.path)
}
