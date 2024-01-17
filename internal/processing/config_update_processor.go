/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package processing

import (
	"errors"
	"fmt"
)

type ConfigUpdateProcessor struct {
	events        map[uint64]*ConfigUpdateEvent
	updateChannel chan *ConfigUpdateEvent
}

func NewConfigUpdateProcessor(updateChannel chan *ConfigUpdateEvent) *ConfigUpdateProcessor {
	return &ConfigUpdateProcessor{
		events:        make(map[uint64]*ConfigUpdateEvent),
		updateChannel: updateChannel,
	}
}

// QueueConfigUpdate adds a ConfigUpdateEvent to the processor's event map.
func (c *ConfigUpdateProcessor) QueueConfigUpdate(config string) (*ConfigUpdateEvent, error) {
	configUpdate := NewConfigUpdateEvent(config)

	if _, found := c.events[configUpdate.EventID]; found {
		errorMessage := fmt.Sprintf("event with key '%d' already exists", configUpdate.EventID)
		configUpdate.SetError(errorMessage)
		return configUpdate, errors.New(errorMessage)
	}

	c.events[configUpdate.EventID] = configUpdate

	c.updateChannel <- configUpdate

	return configUpdate, nil
}

func (c *ConfigUpdateProcessor) LookUpEvent(eventID uint64) (*ConfigUpdateEvent, error) {
	event, found := c.events[eventID]
	if !found {
		return nil, errors.New(fmt.Sprintf("event with id '%d' not found", eventID))
	}
	return event, nil
}
