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

type AppSucces struct {
	Message string      `json:"message"`
	Code    int         `json:"statusCode"`
	Data    interface{} `json:"data"`
}

type AppHandler func(http.ResponseWriter, *http.Request) (*AppSucces, *AppError)

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s, e := fn(w, r)
	if e != nil {
		log.Printf("error: %v", e.Error)
		w.WriteHeader(e.Code)
		jsonResp, err := json.Marshal(e)
		if err != nil {
			log.Printf("error encoding response: %v", err)
			http.Error(w, e.Message, http.StatusInternalServerError)
		}

		w.Write(jsonResp)
	}

	if s != nil {
		w.WriteHeader(s.Code)
		jsonResp, err := json.Marshal(s)
		if err != nil {
			log.Printf("error encoding response: %v", err)
			http.Error(w, e.Message, http.StatusInternalServerError)
		}

		w.Write(jsonResp)
	}
}
