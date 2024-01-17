/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package processing

import (
	json2 "encoding/json"
	"errors"
	"fmt"
	"github.com/cespare/xxhash/v2"
	"time"
)

type EventState int

const (
	Failed EventState = iota
	Accepted
	Started
	Validating
	Validated
	Pending
	Applied
)

func (s EventState) String() string {
	return [...]string{"Failed", "Accepted", "Started", "Validating", "Validated", "Pending", "Applied"}[s]
}

func (s EventState) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

type ConfigUpdateEvent struct {
	ConfigBody string                   `json:"configuration"`
	EventState EventState               `json:"state"`
	EventLog   map[time.Time]EventState `json:"log"`
	Error      string                   `json:"error"`
	EventID    uint64                   `json:"id"`
	Location   string                   `json:"location"`
}

// NewConfigUpdateEvent creates a new ConfigUpdateEvent
//
// [xxhash](https://pkg.go.dev/github.com/cespare/xxhash) is used to generate the hash.
//
// This choice is due to its speed and low collision rate, and the fact that a strong cryptographic hash is not required.
func NewConfigUpdateEvent(configBody string) *ConfigUpdateEvent {
	eventId := xxhash.Sum64String(configBody)
	return &ConfigUpdateEvent{
		ConfigBody: configBody,
		EventState: Accepted,
		EventLog:   map[time.Time]EventState{time.Now(): Accepted},
		EventID:    eventId,
		Location:   fmt.Sprintf("/nginx/config/status/%d", eventId),
	}
}

func (e *ConfigUpdateEvent) SetState(state EventState) error {
	if state == Failed {
		return errors.New("EventState cannot be set to `Failed` directly; use SetError() instead")
	}

	e.EventState = state
	e.EventLog[time.Now()] = state

	return nil
}

func (e *ConfigUpdateEvent) SetError(err string) {
	e.EventState = Failed
	e.Error = err
	e.EventLog[time.Now()] = Failed
}

func (e *ConfigUpdateEvent) Json() ([]byte, error) {
	return json2.Marshal(e)
}
