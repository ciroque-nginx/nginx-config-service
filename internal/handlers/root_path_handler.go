/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"ciroque/go-http-server/internal/metrics"
	"ciroque/go-http-server/internal/responses"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RootPathHandler struct {
	Logger                   *logrus.Entry
	RootPathRequestCount     *prometheus.CounterVec
	RootPathRequestDurations *prometheus.HistogramVec
}

func NewRootPathHandler(logger *logrus.Entry) *RootPathHandler {
	return &RootPathHandler{Logger: logger}
}

func (handler *RootPathHandler) Cleanup() {
	prometheus.Unregister(handler.RootPathRequestCount)
	prometheus.Unregister(handler.RootPathRequestDurations)
}

func (handler *RootPathHandler) BuildHandler() (http.Handler, error) {
	handler.configureMetrics()
	err := handler.registerMetrics()
	if err != nil {
		return nil, err
	}

	h := promhttp.InstrumentHandlerCounter(
		handler.RootPathRequestCount,
		promhttp.InstrumentHandlerDuration(
			handler.RootPathRequestDurations,
			http.HandlerFunc(handler.handler)))

	return h, nil
}

func (handler *RootPathHandler) Route() string {
	return "/"
}

func (handler *RootPathHandler) handler(writer http.ResponseWriter, request *http.Request) {
	handler.Logger.Debugf("Server::ServeRootPath: %#v", request)
	response := responses.NewRootPathResponse("200 OK")

	bytes, err := json.Marshal(&response)
	if err != nil {
		handler.Logger.Warnf("Error responding to request %#v", err)
	}

	writer.Header().Add("Content-Type", "application/json")
	_, err = fmt.Fprintf(writer, "%s", bytes)
	if err != nil {
		handler.Logger.Warnf("Error responding to request %#v", err)
	}

	handler.Logger.Debugf("Server::ServeRootPath: %#v", response)
}

func (handler *RootPathHandler) configureMetrics() {
	handler.RootPathRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metrics.Namespace,
		Subsystem: metrics.RootSubsystem,
		Name:      "root_path_request_count",
		Help:      "The number of requests to the root path",
	}, []string{"code", "method"})

	handler.RootPathRequestDurations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: metrics.Namespace,
		Subsystem: metrics.RootSubsystem,
		Name:      "root_path_request_duration",
		Help:      "Tracks how long it takes to retrieve the root path",
		Buckets:   prometheus.DefBuckets,
	}, []string{"code", "method"})
}

func (handler *RootPathHandler) registerMetrics() error {
	err := prometheus.Register(handler.RootPathRequestCount)
	if err != nil {
		return err
	}

	err = prometheus.Register(handler.RootPathRequestDurations)
	if err != nil {
		return err
	}

	return nil
}
