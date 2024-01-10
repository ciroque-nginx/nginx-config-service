/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace     = "go_http_server"
	RootSubsystem = "root_subsystem"
)

type Metrics struct {
	RootPathRequestCount     *prometheus.CounterVec
	RootPathRequestDurations *prometheus.HistogramVec
}

func NewMetrics() (Metrics, error) {
	metrics := configureMetrics()

	err := prometheus.Register(metrics.RootPathRequestCount)
	if err != nil {
		return metrics, err
	}

	err = prometheus.Register(metrics.RootPathRequestDurations)
	if err != nil {
		return metrics, err
	}

	return metrics, nil
}

func (metrics *Metrics) Shutdown() {
	prometheus.Unregister(metrics.RootPathRequestCount)
	prometheus.Unregister(metrics.RootPathRequestDurations)
}

func configureMetrics() Metrics {
	return Metrics{
		RootPathRequestCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: RootSubsystem,
			Name:      "root_path_request_count",
			Help:      "The number of requests to the root path",
		}, []string{"code", "method"}),
		RootPathRequestDurations: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: RootSubsystem,
			Name:      "root_path_request_duration",
			Help:      "Tracks how long it takes to retrieve the root path",
			Buckets:   prometheus.DefBuckets,
		}, []string{"code", "method"}),
	}
}
