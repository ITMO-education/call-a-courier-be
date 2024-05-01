package rest

import (
	"context"
	"encoding/json"
	"net/http"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/godverv/matreshka"
	"github.com/godverv/matreshka/api"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/itmo-education/delivery-backend/internal/data"
)

type Server struct {
	HttpServer *http.Server

	data data.Data

	version string
}

func NewServer(cfg matreshka.Config, server *api.Rest, db data.Data) *Server {
	r := http.NewServeMux()

	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + server.GetPortStr(),
			Handler: setUpCors().Handler(r),
		},
		data:    db,
		version: cfg.AppInfo().Version,
	}

	r.HandleFunc("/version", s.Version)
	r.HandleFunc("POST /contract", s.RegisterContract)
	r.HandleFunc("GET /contract", s.ListContract)

	return s
}

func (s *Server) Start(_ context.Context) error {
	go func() {
		logrus.Infof("Starting REST server at %s", s.HttpServer.Addr)
		err := s.HttpServer.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err)
		} else {

		}
		logrus.Infof("REST server at %s is stopped", s.HttpServer.Addr)

	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) formResponse(r interface{}) ([]byte, error) {
	return json.Marshal(r)
}

func setUpCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
}
