package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func UnmarshalJSON(r *http.Request, data interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
