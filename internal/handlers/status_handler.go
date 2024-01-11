/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"ciroque/go-http-server/internal/metrics"
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"strings"
)

type StatusHandler struct {
	Context          *Context
	RequestCount     *prometheus.CounterVec
	RequestDurations *prometheus.HistogramVec
}

func NewStatusHandler(context *Context) *StatusHandler {
	return &StatusHandler{Context: context}
}

func (handler *StatusHandler) Cleanup() {
	handler.Context.Logger.Info("RawConfigHandler::Cleanup")
	prometheus.Unregister(handler.RequestCount)
	prometheus.Unregister(handler.RequestDurations)
}

func (handler *StatusHandler) BuildHandler() (http.Handler, error) {
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

func (handler *StatusHandler) Route() string {
	return "/nginx/config/status/"
}

func (handler *StatusHandler) handler(writer http.ResponseWriter, request *http.Request) {
	handler.Context.Logger.Debugf("StatusHandler::handle %#v", request)

	eventId, err := handler.getEventIdFromRequest(request)
	if err != nil {
		WriteErrorResponse(writer, err.Error(), http.StatusBadRequest, handler.Context.Logger)
		return
	}

	updateEvent, err := handler.Context.Processor.LookUpEvent(eventId)
	if err != nil {
		WriteErrorResponse(writer, err.Error(), http.StatusNotFound, handler.Context.Logger)
		return
	}

	WriteUpdateEvent(writer, updateEvent, http.StatusOK, handler.Context.Logger)
}

func (handler *StatusHandler) configureMetrics() {
	handler.RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "status_request_count",
			Help:      "The number of requests served by the status handler",
		}, []string{"code", "method"})

	handler.RequestDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: metrics.RootSubsystem,
			Name:      "status_request_durations",
			Help:      "The duration of requests served by the status handler",
			Buckets:   prometheus.DefBuckets,
		}, []string{"code", "method"})
}

func (handler *StatusHandler) registerMetrics() error {
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

func (handler *StatusHandler) getEventIdFromRequest(request *http.Request) (uint64, error) {
	segments := strings.Split(request.URL.Path, "/")
	if len(segments) < 5 {
		return 0, errors.New("invalid path")
	}

	segment := segments[4]
	handler.Context.Logger.Debugf("StatusHandler::getEventIdFromRequest: %#v", segment)

	eventId, err := strconv.ParseUint(segment, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}

	return eventId, nil
}
