package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var healthcheck Api = Json("GET",
	func(w http.ResponseWriter, r *http.Request, body json.RawMessage) (any, error) {
		return map[string]string{"status": "ok"}, nil
	},
)

func availableLanguages(s *APIServer) Api {
	return Json("GET",
		func(w http.ResponseWriter, r *http.Request, body json.RawMessage) (any, error) {
			return s.runner.AvailableLanguages(), nil
		},
	)
}

func run(s *APIServer) Api {
	return Json("POST",
		func(w http.ResponseWriter, r *http.Request, raw json.RawMessage) (any, error) {
			body := struct {
				Language string   `json:"language"`
				Code     string   `json:"code"`
				Input    []string `json:"input"`
			}{
				Input: []string{""},
			}
			json.Unmarshal(raw, &body)
			if body.Language == "" || body.Code == "" {
				return nil, fmt.Errorf("language and code are required")
			}
			result, err := s.runner.Run(body.Language, body.Code, body.Input)
			if result == nil {
				return nil, err
			}
			if err != nil {
				fmt.Println(err)
			}
			return result, nil
		},
	)
}
