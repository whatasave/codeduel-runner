package api

import (
	"encoding/json"
	"net/http"
)

type Api http.HandlerFunc
type JsonApi func(http.ResponseWriter, *http.Request) (any, error)

func Json(api JsonApi) Api {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if result, err := api(w, r); err != nil {
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
