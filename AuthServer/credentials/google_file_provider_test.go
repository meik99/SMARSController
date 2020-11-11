package credentials

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoogleFileProvider_GetCredentials(t *testing.T) {
	gfp := &googleFileProvider{
		path: "",
		credentialsSourceFn: func(provider *googleFileProvider) io.Reader {
			return strings.NewReader(`{ "web": {"client_id": "test-client"} }`)
		},
	}

	credentials, err := gfp.GetCredentials()

	assert.NoError(t, err)
	assert.NotNil(t, credentials)
	assert.Equal(t, "test-client", credentials.ClientId)
}
