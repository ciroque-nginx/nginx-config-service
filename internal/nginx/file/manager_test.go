/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package file

import (
	"ciroque/go-http-server/internal/config"
	"ciroque/go-http-server/internal/processing"
	context2 "context"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewUpdater(t *testing.T) {
	context := context2.Context(context2.Background())
	updateChannel := make(chan *processing.ConfigUpdateEvent)

	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)

	updaterContext := NewManagerContext(context, loggerEntry, updateChannel)

	settings, _ := config.NewSettings()

	updater := NewConfigFileManager(settings, updaterContext)

	if updater == nil {
		t.Error("Expected updater to not be nil")
	}
}

func TestUpdater_Start(t *testing.T) {
	context := context2.Context(context2.Background())
	updateChannel := make(chan *processing.ConfigUpdateEvent)

	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)

	updaterContext := NewManagerContext(context, loggerEntry, updateChannel)

	settings, _ := config.NewSettings()

	updater := NewConfigFileManager(settings, updaterContext)

	go updater.Start()

	updateEvent := processing.NewConfigUpdateEvent("Config Body")

	updateChannel <- updateEvent
}

func TestUpdater_StartClosedChannel(t *testing.T) {
	context := context2.Context(context2.Background())
	updateChannel := make(chan *processing.ConfigUpdateEvent)

	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)

	updaterContext := NewManagerContext(context, loggerEntry, updateChannel)

	settings, _ := config.NewSettings()

	updater := NewConfigFileManager(settings, updaterContext)

	go updater.Start()

	close(updateChannel)
}

func TestUpdater_StartContextDone(t *testing.T) {
	context, cancel := context2.WithCancel(context2.Background())
	updateChannel := make(chan *processing.ConfigUpdateEvent)

	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)

	updaterContext := NewManagerContext(context, loggerEntry, updateChannel)

	settings, _ := config.NewSettings()

	updater := NewConfigFileManager(settings, updaterContext)

	go updater.Start()

	cancel()
}

func TestUpdater_StartContextDoneAfterWork(t *testing.T) {
	context, cancel := context2.WithCancel(context2.Background())
	updateChannel := make(chan *processing.ConfigUpdateEvent)

	logger := logrus.New()
	loggerEntry := logrus.NewEntry(logger)

	updaterContext := NewManagerContext(context, loggerEntry, updateChannel)

	settings, _ := config.NewSettings()

	updater := NewConfigFileManager(settings, updaterContext)

	go updater.Start()

	updateEvent := processing.NewConfigUpdateEvent("Config Body")

	updateChannel <- updateEvent

	cancel()

	time.Sleep(2 * time.Second)
}
