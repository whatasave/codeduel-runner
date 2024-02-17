package api

import "net/http"

var healthcheck Api = Json(
	func(w http.ResponseWriter, r *http.Request) (any, error) {
		return map[string]string{"status": "ok"}, nil
	},
)

func run(s *APIServer) Api {
	return Json(
		func(w http.ResponseWriter, r *http.Request) (any, error) {
			return map[string]string{"status": "ok"}, nil
		},
	)
}
