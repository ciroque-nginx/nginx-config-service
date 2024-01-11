/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package processing

import (
	"fmt"
	"testing"
)

// TODO: Provide a channel that works with the test
func TestConfigUpdateProcessor_QueueConfigUpdate(t *testing.T) {
	configBody := "Sample NGINX Configuration"
	updateChannel := make(chan *ConfigUpdateEvent)

	go func() {
		for {
			select {
			case updateEvent := <-updateChannel:
				fmt.Printf("updateEvent: %v\n", updateEvent)
			}
		}
	}()

	processor := NewConfigUpdateProcessor(updateChannel)
	configUpdateEvent, err := processor.QueueConfigUpdate(configBody)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if configUpdateEvent.EventID == 0 {
		t.Errorf("expected eventID to be initialized")
	}
}

func TestConfigUpdateProcessor_QueueConfigUpdate_Duplicate(t *testing.T) {
	configBody := "Sample NGINX Configuration"
	updateChannel := make(chan *ConfigUpdateEvent)

	go func() {
		for {
			select {
			case updateEvent := <-updateChannel:
				fmt.Printf("updateEvent: %v\n", updateEvent)
			}
		}
	}()

	processor := NewConfigUpdateProcessor(updateChannel)

	configUpdateEvent, err := processor.QueueConfigUpdate(configBody)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if configUpdateEvent.EventID == 0 {
		t.Errorf("expected eventID to be intialized")
	}

	duplicateUpdateEvent, err := processor.QueueConfigUpdate(configBody)
	if err == nil {
		t.Errorf("expected error, got none")
	}

	// The found id is returned
	if duplicateUpdateEvent.EventID != configUpdateEvent.EventID {
		t.Errorf("expected nextEventID to be uninitialized, got %d", duplicateUpdateEvent.EventID)
	}

	if err.Error() != fmt.Sprintf("event with key '%d' already exists", configUpdateEvent.EventID) {
		t.Errorf("expected error to be 'event with key '%d' already exists', got %s", configUpdateEvent.EventID, err.Error())
	}
}

func TestConfigUpdateProcessor_LookUpEvent(t *testing.T) {
	configBody := "Sample NGINX Configuration"
	updateChannel := make(chan *ConfigUpdateEvent)

	go func() {
		for {
			select {
			case updateEvent := <-updateChannel:
				fmt.Printf("updateEvent: %v\n", updateEvent)
			}
		}
	}()

	processor := NewConfigUpdateProcessor(updateChannel)

	updateEvent, err := processor.QueueConfigUpdate(configBody)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if updateEvent.EventID == 0 {
		t.Errorf("expected eventID to be intialized")
	}

	event, err := processor.LookUpEvent(updateEvent.EventID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if event == nil {
		t.Errorf("expected event to be non-nil")
	}

	if event.ConfigBody != string(configBody) {
		t.Errorf("expected event.ConfigurationBody to be 'test', got %s", event.ConfigBody)
	}
}

func TestConfigUpdateProcessor_LookUpEvent_NotFound(t *testing.T) {
	eventId := uint64(1)
	updateChannel := make(chan *ConfigUpdateEvent)
	processor := NewConfigUpdateProcessor(updateChannel)

	event, err := processor.LookUpEvent(eventId)
	if err == nil {
		t.Errorf("expected error, got none")
	}

	if event != nil {
		t.Errorf("expected event to be nil")
	}

	if err.Error() != fmt.Sprintf("event with id '%d' not found", eventId) {
		t.Errorf("expected error to be 'event with id '%d' not found', got %s", eventId, err.Error())
	}
}
