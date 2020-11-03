package main

type googleCredentials struct {
	Web webCredentials `json:"web"`
}

type webCredentials struct {
	ClientId            string `json:"client_id"`
	ProjectId           string `json:"project_id"`
	AuthURI             string `json:"auth_uri"`
	TokenURI            string `json:"token_uri"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
	ClientSecret        string `json:"client_secret"`
}
