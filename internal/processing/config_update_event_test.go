/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package processing

import (
	"strings"
	"testing"
)

func TestNewConfigUpdateEvent(t *testing.T) {
	configurationBody := "Text NGINX Configuration"
	event := NewConfigUpdateEvent(configurationBody)

	if event == nil {
		t.Errorf("Expected event to not be nil")
	}

	if event.ConfigBody != configurationBody {
		t.Errorf("Expected event.ConfigurationBody to be '%s', but was '%s'", configurationBody, event.ConfigBody)
	}

	if event.EventState != Accepted {
		t.Errorf("Expected event.EventState to be 'Accepted', but was '%d'", event.EventState)
	}

	if event.Error != "" {
		t.Errorf("Expected event.Error to be nil, but was '%s'", event.Error)
	}
}

func TestConfigUpdateEvent_SetState(t *testing.T) {
	configurationBody := "Text NGINX Configuration"
	event := NewConfigUpdateEvent(configurationBody)

	err := event.SetState(Started)
	if err != nil {
		t.Errorf("Expected event.SetState() to not return an error, but got '%s'", err)
	}

	if event.EventState != Started {
		t.Errorf("Expected event.EventState to be 'Started', but was '%d'", event.EventState)
	}

	if event.Error != "" {
		t.Errorf("Expected event.Error to be nil, but was '%s'", event.Error)
	}
}

func TestConfigUpdateEvent_SetError(t *testing.T) {
	errorMessage := "Something went horribly, horribly wrong"
	configurationBody := "Text NGINX Configuration"
	event := NewConfigUpdateEvent(configurationBody)

	event.SetError(errorMessage)

	if event.EventState != Failed {
		t.Errorf("Expected event.EventState to be 'Failed', but was '%d'", event.EventState)
	}

	if event.Error != "" && event.Error != errorMessage {
		t.Errorf("Expected event.Error to be '%s', but was '%s'", errorMessage, event.Error)
	}
}

func TestConfigUpdateEvent_SetErrorFail(t *testing.T) {
	configurationBody := "Test NGINX Configuration"
	event := NewConfigUpdateEvent(configurationBody)

	err := event.SetState(Failed)
	if err == nil {
		t.Errorf("Expected event.SetState() to return an error, but got nil")
	}

	if event.EventState != Accepted {
		t.Errorf("Expected event.EventState to be 'Accepted', but was '%d'", event.EventState)
	}

	if event.Error != "" {
		t.Errorf("Expected event.Error to be nil, but was '%s'", event.Error)
	}

	if err != nil && err.Error() != "EventState cannot be set to `Failed` directly; use SetError() instead" {
		t.Errorf("Expected event.SetState() to return an error, but got '%s'", err)
	}
}

func TestEventState_MarshalJSON(t *testing.T) {
	state := Accepted
	json, err := state.MarshalJSON()
	if err != nil {
		t.Errorf("Expected MarshalJSON() to not return an error, but got '%s'", err)
	}

	if string(json) != `"Accepted"` {
		t.Errorf("Expected MarshalJSON() to return 'Accepted', but got '%s'", string(json))
	}
}

func TestEventState_String(t *testing.T) {
	state := Accepted
	if state.String() != "Accepted" {
		t.Errorf("Expected String() to return 'Accepted', but got '%s'", state.String())
	}
}

func TestConfigUpdateEvent_Json(t *testing.T) {
	configurationBody := "Text NGINX Configuration"
	event := NewConfigUpdateEvent(configurationBody)

	json, err := event.Json()
	if err != nil {
		t.Errorf("Expected Json() to not return an error, but got '%s'", err)
	}

	if !strings.Contains(string(json), `Text NGINX Configuration`) && strings.Contains(string(json), `Accepted`) {
		t.Errorf("Expected Json() to return '{\"configuration\": \"Text NGINX Configuration\", \"state\": \"Accepted\"}', but got '%s'", string(json))
	}
}
