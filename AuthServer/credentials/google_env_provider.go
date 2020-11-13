package credentials

import (
	"github.com/meik99/CoffeeToGO/AuthServer/auth_json"
	"os"
	"strings"
)

type googleEnvironmentProvider struct {
	envVarName string
}

func NewGoogleEnvironmentProvider(envVarName string) Provider {
	return &googleEnvironmentProvider{
		envVarName: envVarName,
	}
}

func (gep *googleEnvironmentProvider) GetCredentials() (OAuthCredentials, error) {
	var gCredentials googleCredentials
	err := auth_json.ParseJSONToInterface(strings.NewReader(os.Getenv(gep.envVarName)), &gCredentials)
	if err != nil {
		return OAuthCredentials{}, err
	}
	return gCredentials.Web, nil
}
