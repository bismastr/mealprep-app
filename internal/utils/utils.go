package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

func GetIntFromValue(r *http.Request, key string, defaultValue int) (int, error) {
	valueStr := r.FormValue(key)

	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(valueStr)
}
