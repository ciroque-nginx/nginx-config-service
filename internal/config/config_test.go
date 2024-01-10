/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestNewSettings(t *testing.T) {
	config, err := NewSettings()
	if err != nil {
		t.Errorf("NewSettings() returned an error: %v", err)
	}

	if config.Port != DefaultPort {
		t.Errorf("NewSettings() returned an unexpected port. Expected: %v, Actual: %v", DefaultPort, config.Port)
	}

	if config.Host != DefaultHost {
		t.Errorf("NewSettings() returned an unexpected host. Expected: %v, Actual: %v", DefaultHost, config.Host)
	}

	if config.LogLevel != DefaultLogLevel {
		t.Errorf("NewSettings() returned an unexpected log level. Expected: %v, Actual: %v", DefaultLogLevel, config.LogLevel)
	}
}

func TestOverrideDefaultPort(t *testing.T) {
	expectedPort := "9999"
	err := os.Setenv(PortEnvironmentVariable, expectedPort)
	if err != nil {
		t.Errorf("Failed to set environment variable %v: %v", PortEnvironmentVariable, err)
	}

	config, err := NewSettings()
	if err != nil {
		t.Errorf("NewSettings() returned an error: %v", err)
	}

	if config.Port != expectedPort {
		t.Errorf("NewSettings() returned an unexpected port. Expected: %v, Actual: %v", expectedPort, config.Port)
	}
}

func TestOverrideDefaultHost(t *testing.T) {
	expectedHost := "127.0.0.1"
	err := os.Setenv(HostEnvironmentVariable, expectedHost)
	if err != nil {
		t.Errorf("Failed to set environment variable %v: %v", HostEnvironmentVariable, err)
	}

	config, err := NewSettings()
	if err != nil {
		t.Errorf("NewSettings() returned an error: %v", err)
	}

	if config.Host != expectedHost {
		t.Errorf("NewSettings() returned an unexpected host. Expected: %v, Actual: %v", expectedHost, config.Host)
	}
}

func TestOverrideDefaultLogLevel_Debug(t *testing.T) {

	logLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}

	for _, logLevel := range logLevels {
		expectedLogLevel := logLevel.String()
		err := os.Setenv(LogLevelEnvironmentVariable, expectedLogLevel)
		if err != nil {
			t.Errorf("Failed to set environment variable %v: %v", LogLevelEnvironmentVariable, err)
		}

		config, err := NewSettings()
		if err != nil {
			t.Errorf("NewSettings() returned an error: %v", err)
		}

		if config.LogLevel != logLevel {
			t.Errorf("NewSettings() returned an unexpected log level. Expected: %v, Actual: %v", logLevel, config.LogLevel)
		}
	}
}
