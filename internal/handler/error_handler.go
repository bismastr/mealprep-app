package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type AppError struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"statusCode"`
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		log.Printf("error: %v", e.Error)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(e.Code)

		jsonResp, err := json.Marshal(e)
		if err != nil {
			log.Printf("error encoding response: %v", err)
			http.Error(w, e.Message, http.StatusInternalServerError)
		}

		w.Write(jsonResp)
	}
}
