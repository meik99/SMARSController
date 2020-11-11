package credentials

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Provider interface {
	GetCredentials() (OAuthCredentials, error)
}

type googleCredentials struct {
	Web OAuthCredentials `json:"web"`
}

type OAuthCredentials struct {
	ClientId            string `json:"client_id"`
	ProjectId           string `json:"project_id"`
	AuthURI             string `json:"auth_uri"`
	TokenURI            string `json:"token_uri"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
	ClientSecret        string `json:"client_secret"`
}

func fromReader(reader io.Reader, v interface{}) error {
	credentialsJsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(credentialsJsonData, &v)
	return err
}
