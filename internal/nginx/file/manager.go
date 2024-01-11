/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package file

import (
	"ciroque/go-http-server/internal/config"
	"ciroque/go-http-server/internal/processing"
)

type ConfigFileManager struct {
	settings *config.Settings
	context  *ManagerContext
}

func NewConfigFileManager(settings *config.Settings, context *ManagerContext) *ConfigFileManager {
	return &ConfigFileManager{
		settings: settings,
		context:  context,
	}
}

func (manager *ConfigFileManager) Start() {
	for {
		select {
		case event, open := <-manager.context.UpdateChannel:
			if !open {
				manager.context.Logger.Info("Update channel closed")
				return
			}

			err := processUpdateEvent(event)
			manager.context.Logger.Infof("Processed event: %v\n", event)
			if err != nil {
				manager.context.Logger.Errorf("Error processing event: %v\n", err)
			}

			select {
			case <-manager.context.Context.Done():
				manager.context.Logger.Info("ManagerContext done")
				return
			}

		case <-manager.context.Context.Done():
			manager.context.Logger.Info("ManagerContext done")
			return
		}
	}
}

func processUpdateEvent(event *processing.ConfigUpdateEvent) error {

	if event.EventID/3 == 0 {
		event.SetError("This is an error")
		return nil
	}

	event.SetState(processing.Started)

	event.SetState(processing.Validating)

	event.SetState(processing.Validated)

	event.SetState(processing.Pending)

	event.SetState(processing.Applied)

	return nil
}
