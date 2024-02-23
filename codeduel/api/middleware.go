package api

import (
	"encoding/json"
	"net/http"
)

type Api http.HandlerFunc
type JsonApi func(http.ResponseWriter, *http.Request, json.RawMessage) (any, error)

func Json(method string, api JsonApi) Api {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 page not found"))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		var body json.RawMessage = []byte{}
		if r.Method != "GET" {
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ApiError{Message: "Body should be a valid JSON object"})
				return
			}
		}
		if result, err := api(w, r, body); err != nil {
			if w.Header().Get("status") == "" {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
		} else {
			if w.Header().Get("status") == "" {
				w.WriteHeader(http.StatusOK)
			}
			json.NewEncoder(w).Encode(result)
		}
	}
}
