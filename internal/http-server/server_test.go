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
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

type TestHandler struct {
}

func (handler *TestHandler) BuildHandler() (http.Handler, error) {
	return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})), nil
}

func (handler *TestHandler) Cleanup() {

}

func (handler *TestHandler) Route() string {
	return "/"
}

func TestServer_Run(t *testing.T) {
	port := rand.Intn(1000) + 20000

	abortChannel := make(chan string)
	defer close(abortChannel)

	server, err := buildServer(abortChannel, port)
	if err != nil {
		t.Fatalf("Error building server: %v", err)
	}

	stopCh := make(chan struct{})

	go server.Run(stopCh)

	select {
	case err := <-abortChannel:
		{
			t.Fatalf("Error starting server: %v", err)
		}
	case <-time.After(3 * time.Second):
		{
			close(stopCh)
		}
	}
}

func buildServer(abortChannel chan string, port int) (*Server, error) {
	os.Setenv("NCS_HTTP_PORT", fmt.Sprintf("%d", port))

	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("Error creating configuration settings: %v", err)
	}

	logger := logrus.New()
	logger.SetLevel(settings.LogLevel)

	routes := []routehandlers.Interface{
		&TestHandler{},
	}

	server, err := NewServer(
		settings,
		abortChannel,
		logrus.NewEntry(logger),
		routes)

	if err != nil {
		return nil, err
	}

	return server, err
}
