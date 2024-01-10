/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package httpserver

import (
	"ciroque/go-http-server/internal/config"
	"ciroque/go-http-server/internal/metrics"
	"ciroque/go-http-server/internal/responses"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	AbortChannel chan<- string
	Logger       *logrus.Entry
	Settings     *config.Settings
	Metrics      *metrics.Metrics
}

func NewServer(
	settings *config.Settings,
	abortChannel chan<- string,
	logger *logrus.Entry,
	metricsClient *metrics.Metrics) (*Server, error) {

	server := Server{
		AbortChannel: abortChannel,
		Logger:       logger,
		Settings:     settings,
		Metrics:      metricsClient,
	}

	return &server, nil
}

func (server *Server) Run(stopCh <-chan struct{}) {
	server.Logger.Debugf("Server::Run: %#v", server)
	rootPathHandler := promhttp.InstrumentHandlerCounter(
		server.Metrics.RootPathRequestCount,
		promhttp.InstrumentHandlerDuration(
			server.Metrics.RootPathRequestDurations,
			http.HandlerFunc(server.ServeRootPath)))

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", rootPathHandler)

	address := fmt.Sprintf("%s:%s", server.Settings.Host, server.Settings.Port)

	server.Logger.Info("Listening on ", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		server.AbortChannel <- err.Error()
	}

	<-stopCh

	server.Logger.Info("Server stopped")
}

func (server *Server) ServeRootPath(writer http.ResponseWriter, request *http.Request) {
	server.Logger.Debugf("Server::ServeRootPath: %#v", request)
	response := responses.NewRootPathResponse("200 OK")

	bytes, err := json.Marshal(&response)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}

	writer.Header().Add("Content-Type", "application/json")
	_, err = fmt.Fprintf(writer, "%s", bytes)
	if err != nil {
		server.Logger.Warnf("Error responding to request %#v", err)
	}

	server.Logger.Debugf("Server::ServeRootPath: %#v", response)
}
