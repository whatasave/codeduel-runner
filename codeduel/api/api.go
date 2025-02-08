package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var handlerHealthCheck Api = Json("GET",
	func(w http.ResponseWriter, r *http.Request, body json.RawMessage) (any, error) {
		return map[string]string{"status": "ok"}, nil
	},
)

func handlerGetAvailableLanguages(apiServer *APIServer) Api {
	return Json("GET",
		func(w http.ResponseWriter, r *http.Request, body json.RawMessage) (any, error) {
			return apiServer.runner.AvailableLanguages(), nil
		},
	)
}

func handlePostRunCode(apiServer *APIServer) Api {
	return Json("POST",
		func(w http.ResponseWriter, r *http.Request, raw json.RawMessage) (any, error) {
			body := struct {
				Language string   `json:"language"`
				Code     string   `json:"code"`
				Input    []string `json:"input"` // optional
			}{
				Input: []string{""},
			}

			json.Unmarshal(raw, &body)
			if body.Language == "" || body.Code == "" {
				return nil, fmt.Errorf("language and code are required")
			}
			result, err := apiServer.runner.Run(body.Language, body.Code, body.Input)
			if result == nil {
				return nil, err
			}
			if err != nil {
				log.Println(err)
			}
			return result, nil
		},
	)
}
