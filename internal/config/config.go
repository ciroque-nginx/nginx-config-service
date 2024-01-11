/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	DefaultUpdateChannelSize    = 100
	DefaultHost                 = "0.0.0.0"
	DefaultPort                 = "8293"
	DefaultLogLevel             = logrus.InfoLevel
	HostEnvironmentVariable     = "NCS_HTTP_HOST"
	LogLevelEnvironmentVariable = "NCS_LOG_LEVEL"
	PortEnvironmentVariable     = "NCS_HTTP_PORT"
)

type Settings struct {
	UpdateChannelSize int
	Host              string
	Port              string
	LogLevel          logrus.Level
}

func NewSettings() (*Settings, error) {
	port := os.Getenv(PortEnvironmentVariable)
	if port == "" {
		port = DefaultPort
	}

	host := os.Getenv(HostEnvironmentVariable)
	if host == "" {
		host = DefaultHost
	}

	logLevel := getLogLevel()

	config := &Settings{DefaultUpdateChannelSize, host, port, logLevel}

	return config, nil
}

func getLogLevel() logrus.Level {
	logLevel := os.Getenv(LogLevelEnvironmentVariable)
	logrus.Debugf("Settings::setLogLevel: %s", logLevel)
	switch logLevel {
	case "panic":
		return logrus.PanicLevel

	case "fatal":
		return logrus.FatalLevel

	case "error":
		return logrus.ErrorLevel

	case "warning":
		return logrus.WarnLevel

	case "info":
		return logrus.InfoLevel

	case "debug":
		return logrus.DebugLevel

	case "trace":
		return logrus.TraceLevel

	default:
		return DefaultLogLevel
	}
}
