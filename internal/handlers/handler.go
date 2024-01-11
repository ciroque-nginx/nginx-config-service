/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Interface interface {
	BuildHandler() (http.Handler, error)
	Cleanup()
	Route() string
}

func GetRoutes(logger *logrus.Entry) []Interface {
	return []Interface{
		NewRootPathHandler(logger),
		NewMetricsHandler(),
	}
}
