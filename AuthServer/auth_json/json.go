package auth_json

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseJSONToInterface(reader io.Reader, v interface{}) error {
	credentialsJsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(credentialsJsonData, &v)
	return err
}
