package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/xedom/codeduel/codeduel/runner"
)

type APIServer struct {
	host   string
	port   string
	runner *runner.Runner
}

type ApiError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type ApiResult struct {
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

func NewAPIServer(host, port string) (*APIServer, error) {
	runner, err := runner.NewRunner()
	if err != nil {
		return nil, err
	}

	return &APIServer{
		host:   host,
		port:   port,
		runner: runner,
	}, nil
}

func (s *APIServer) Run() {
	log.Printf("[API] Starting API server on http://%s:%s", s.host, s.port)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1", healthcheck)
	router.HandleFunc("/api/v1/run", run(s))
	router.HandleFunc("/api/v1/languages", availableLanguages(s))

	frontendUrl := os.Getenv("FRONTEND_URL")

	address := fmt.Sprintf("%s:%s", s.host, s.port)
	err := http.ListenAndServe(address, handlers.CORS(
		handlers.AllowedOrigins([]string{frontendUrl}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "x-token"}),
		handlers.AllowCredentials(),
	)(router))

	if err != nil {
		log.Fatal("[API] Error starting API server: ", err)
	}
}
