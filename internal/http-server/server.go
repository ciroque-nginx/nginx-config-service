/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package httpserver

import (
	"ciroque/go-http-server/internal/config"
	routehandlers "ciroque/go-http-server/internal/handlers"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	AbortChannel chan<- string
	Logger       *logrus.Entry
	Settings     *config.Settings
	Routes       []routehandlers.Interface
}

func NewServer(
	settings *config.Settings,
	abortChannel chan<- string,
	logger *logrus.Entry,
	routes []routehandlers.Interface) (*Server, error) {

	server := Server{
		AbortChannel: abortChannel,
		Logger:       logger,
		Settings:     settings,
		Routes:       routes,
	}

	return &server, nil
}

func (server *Server) Run(stopCh <-chan struct{}) {
	server.Logger.Debugf("Server::Run: %#v", server)

	mux := http.NewServeMux()

	for _, route := range server.Routes {
		handler, err := route.BuildHandler()
		if err != nil {
			server.Logger.Fatalf("Error building handler %#v", err)
		}

		mux.Handle(route.Route(), handler)

		defer route.Cleanup()
	}

	address := fmt.Sprintf("%s:%s", server.Settings.Host, server.Settings.Port)

	server.Logger.Info("Listening on ", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		server.AbortChannel <- err.Error()
	}

	<-stopCh

	server.Logger.Info("Server stopped")
}
