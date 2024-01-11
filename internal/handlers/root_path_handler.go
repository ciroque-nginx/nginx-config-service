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
	"net/http"
)

type RootPathHandler struct {
	Context          *Context
	RequestCount     *prometheus.CounterVec
	RequestDurations *prometheus.HistogramVec
}

func NewRootPathHandler(context *Context) *RootPathHandler {
	return &RootPathHandler{context, nil, nil}
}

func (handler *RootPathHandler) Cleanup() {
	handler.Context.Logger.Info("RootPathHandler::Cleanup")
	prometheus.Unregister(handler.RequestCount)
	prometheus.Unregister(handler.RequestDurations)
}

func (handler *RootPathHandler) BuildHandler() (http.Handler, error) {
	handler.configureMetrics()
	err := handler.registerMetrics()
	if err != nil {
		return nil, err
	}

	h := promhttp.InstrumentHandlerCounter(
		handler.RequestCount,
		promhttp.InstrumentHandlerDuration(
			handler.RequestDurations,
			http.HandlerFunc(handler.handler)))

	return h, nil
}

func (handler *RootPathHandler) Route() string {
	return "/"
}

func (handler *RootPathHandler) handler(writer http.ResponseWriter, request *http.Request) {
	handler.Context.Logger.Debugf("Server::ServeRootPath: %#v", request)
	response := responses.NewJsonBodyResponse("200 OK")

	bytes, err := json.Marshal(&response)
	if err != nil {
		handler.Context.Logger.Warnf("Error responding to request %#v", err)
	}

	writer.Header().Add("Content-Type", "application/json")
	_, err = fmt.Fprintf(writer, "%s", bytes)
	if err != nil {
		handler.Context.Logger.Warnf("Error responding to request %#v", err)
	}

	handler.Context.Logger.Debugf("Server::ServeRootPath: %#v", response)
}

func (handler *RootPathHandler) configureMetrics() {
	handler.RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "root_path_request_count",
			Help:      "The number of requests to the root path",
		}, []string{"code", "method"})

	handler.RequestDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "root_path_request_duration",
			Help:      "Tracks how long it takes to retrieve the root path",
			Buckets:   prometheus.DefBuckets,
		}, []string{"code", "method"})
}

func (handler *RootPathHandler) registerMetrics() error {
	err := prometheus.Register(handler.RequestCount)
	if err != nil {
		return err
	}

	err = prometheus.Register(handler.RequestDurations)
	if err != nil {
		return err
	}

	return nil
}
