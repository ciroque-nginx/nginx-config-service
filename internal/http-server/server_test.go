/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package httpserver

import (
	"ciroque/go-http-server/internal/config"
	"ciroque/go-http-server/internal/metrics"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer_Run(t *testing.T) {
	abortChannel := make(chan string)
	defer close(abortChannel)

	server, err := buildServer(abortChannel)
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

func TestServer_AddressAlreadyInUse(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)

	abortChannelOne := make(chan string)
	//defer close(abortChannelOne)

	abortChannelTwo := make(chan string)
	defer close(abortChannelTwo)

	serverOne, err := buildServer(abortChannelOne)
	if err != nil {
		t.Fatalf("Error building serverOne: %v", err)
	}
	go serverOne.Run(stopCh)

	serverTwo, err := buildServer(abortChannelTwo)
	if err != nil {
		t.Fatalf("Error building serverTwo: %v", err)
	}
	go serverTwo.Run(stopCh)

	select {
	case err := <-abortChannelTwo:
		{
			if err != "listen tcp 0.0.0.0:8293: bind: address already in use" {
				t.Fatalf(">>>> Error starting serverTwo: %v", err)
			}
		}
	case <-time.After(3 * time.Second):
		{
			t.Errorf("Server did not fail to start")
		}
	}
}

func TestServer_ServeRootPath(t *testing.T) {
	abortChannel := make(chan string)
	defer close(abortChannel)

	server, err := buildServer(abortChannel)
	if err != nil {
		t.Fatalf("Error building server: %v", err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	rr := httptest.NewRecorder()

	server.ServeRootPath(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func buildServer(abortChannel chan string) (*Server, error) {
	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("Error creating configuration settings: %v", err)
	}

	logger := logrus.New()
	logger.SetLevel(settings.LogLevel)

	metricClient, err := metrics.NewMetrics()
	if err != nil {
		logger.Fatalf("Error creating metrics client: %v", err)
	}
	defer metricClient.Shutdown()

	server, err := NewServer(
		settings,
		abortChannel,
		logrus.NewEntry(logger),
		&metricClient)

	if err != nil {
		return nil, err
	}

	return server, err
}
