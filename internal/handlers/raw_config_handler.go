/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"ciroque/go-http-server/internal/metrics"
	"ciroque/go-http-server/internal/nginx/configs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
)

type RawConfigHandler struct {
	Context          *Context
	RequestCount     *prometheus.CounterVec
	RequestDurations *prometheus.HistogramVec
}

func NewRawConfigHandler(context *Context) *RawConfigHandler {
	return &RawConfigHandler{Context: context}
}

func (handler *RawConfigHandler) Cleanup() {
	handler.Context.Logger.Info("RawConfigHandler::Cleanup")
	prometheus.Unregister(handler.RequestCount)
	prometheus.Unregister(handler.RequestDurations)
}

func (handler *RawConfigHandler) BuildHandler() (http.Handler, error) {
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

func (handler *RawConfigHandler) Route() string {
	return "/nginx/config"
}

func (handler *RawConfigHandler) handler(writer http.ResponseWriter, request *http.Request) {
	responseStatusCode := http.StatusAccepted

	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		WriteErrorResponse(writer, "Error reading request body", http.StatusInternalServerError, handler.Context.Logger)
		return
	}

	if len(body) == 0 {
		WriteErrorResponse(writer, "An NGINX configuration was not found in the request", http.StatusBadRequest, handler.Context.Logger)
		return
	}

	rawConfig, err := configs.NewRawConfig(body).GetConfig()
	if err != nil {
		WriteErrorResponse(writer, "Error parsing the NGINX configuration", http.StatusBadRequest, handler.Context.Logger)
		return
	}

	configUpdateEvent, err := handler.Context.Processor.QueueConfigUpdate(rawConfig)
	if err != nil {
		handler.Context.Logger.Warnf("Error queuing the NGINX configuration, %#v", err)
		responseStatusCode = http.StatusConflict
	}

	writer.Header().Add("Location", configUpdateEvent.Location)

	// Close the request body to release resources
	defer request.Body.Close()

	WriteUpdateEvent(writer, configUpdateEvent, responseStatusCode, handler.Context.Logger)
}

func (handler *RawConfigHandler) configureMetrics() {
	handler.RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "nginx_config_post_request_count",
			Help:      "The number of requests served by the nginx config post handler",
		}, []string{"code", "method"})

	handler.RequestDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "nginx_config_post_request_durations",
			Help:      "The duration of requests served by the nginx config post handler",
			Buckets:   prometheus.DefBuckets,
		}, []string{"code", "method"})
}

func (handler *RawConfigHandler) registerMetrics() error {
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
