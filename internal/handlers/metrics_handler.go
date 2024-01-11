/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type MetricsHandler struct {
}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{}
}

func (handler *MetricsHandler) Cleanup() {
}

func (handler *MetricsHandler) Route() string {
	return "/metrics"
}

func (handler *MetricsHandler) BuildHandler() (http.Handler, error) {
	return promhttp.Handler(), nil
}
