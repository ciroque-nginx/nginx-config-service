/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package handlers

import (
	"ciroque/go-http-server/internal/processing"
	"ciroque/go-http-server/internal/responses"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Interface interface {
	BuildHandler() (http.Handler, error)
	Cleanup()
	Route() string
}

func GetRoutes(context *Context) []Interface {
	return []Interface{
		NewMetricsHandler(),
		NewRawConfigHandler(context),
		NewStatusHandler(context),
		NewRootPathHandler(context),
	}
}

func WriteErrorResponse(writer http.ResponseWriter, message string, statusCode int, logger *logrus.Entry) {
	response := responses.NewJsonBodyResponse(message)
	bytes, err := json.Marshal(response)
	if err != nil {
		correlationId := uuid.NewString()
		logger.Errorf("(%#v) Error marshalling response %#v", correlationId, err)
		statusCode = http.StatusInternalServerError
		bytes = []byte(fmt.Sprintf("Error marshalling response, provide this correlation id when reporting the problem: %#v", correlationId))
	}

	writer.Header().Add("Content-Type", "application/json")
	http.Error(writer, string(bytes), statusCode)
}

func WriteUpdateEvent(writer http.ResponseWriter, updateEvent *processing.ConfigUpdateEvent, statusCode int, logger *logrus.Entry) {
	responseBody, err := updateEvent.Json()
	if err != nil {
		correlationId := uuid.NewString()
		logger.Errorf("(%#v) Error marshalling response %#v", correlationId, err)
		statusCode = http.StatusInternalServerError
		responseBody = []byte(fmt.Sprintf("Error marshalling response, provide this correlation id when reporting the problem: %#v", correlationId))
	}

	writer.WriteHeader(statusCode)

	_, err = writer.Write(responseBody)
	if err != nil {
		logger.Errorf("Error writing response %#v", err)
		return
	}
}
