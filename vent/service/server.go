// Copyright 2019 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"context"
	"net/http"

	"github.com/hyperledger/burrow/vent/config"
	"github.com/hyperledger/burrow/vent/logger"
)

// Server exposes HTTP endpoints for the service
type Server struct {
	Config   *config.VentConfig
	Log      *logger.Logger
	Consumer *Consumer
	mux      *http.ServeMux
	stopCh   chan bool
}

// NewServer returns a new HTTP server
func NewServer(cfg *config.VentConfig, log *logger.Logger, consumer *Consumer) *Server {
	// setup handlers
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler(consumer))

	return &Server{
		Config:   cfg,
		Log:      log,
		Consumer: consumer,
		mux:      mux,
		stopCh:   make(chan bool, 1),
	}
}

// Run starts the HTTP server
func (s *Server) Run() {
	s.Log.Info("msg", "Starting HTTP Server")

	// start http server
	httpServer := &http.Server{Addr: s.Config.HTTPAddr, Handler: s}

	go func() {
		s.Log.Info("msg", "HTTP Server listening", "address", s.Config.HTTPAddr)
		httpServer.ListenAndServe()
	}()

	// wait for stop signal
	<-s.stopCh

	s.Log.Info("msg", "Shutting down HTTP Server...")

	httpServer.Shutdown(context.Background())
}

// ServeHTTP dispatches the HTTP requests using the Server Mux
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(resp, req)
}

// Shutdown gracefully shuts down the HTTP Server
func (s *Server) Shutdown() {
	s.stopCh <- true
}

func healthHandler(consumer *Consumer) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		err := consumer.Health()
		if err != nil {
			resp.WriteHeader(http.StatusServiceUnavailable)
		} else {
			resp.WriteHeader(http.StatusOK)
		}
	}
}
