package credentials

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGoogleEnvironmentProvider_GetCredentials(t *testing.T) {
	envVarName := "COFFEE_AUTH_TEST_CREDENTIALS"
	err := os.Setenv(envVarName, `{ "web" : {"client_id": "test-client" } }`)

	assert.NoError(t, err)

	gep := NewGoogleEnvironmentProvider(envVarName)
	credentials, err := gep.GetCredentials()

	assert.NoError(t, err)
	assert.NotNil(t, credentials)
	assert.Equal(t, "test-client", credentials.ClientId)

	err = os.Unsetenv(envVarName)

	assert.NoError(t, err)
}
