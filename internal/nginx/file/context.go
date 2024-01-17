/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package file

import (
	"ciroque/go-http-server/internal/processing"
	"context"
	"github.com/sirupsen/logrus"
)

type ManagerContext struct {
	Context       context.Context
	Logger        *logrus.Entry
	UpdateChannel chan *processing.ConfigUpdateEvent
}

func NewManagerContext(context context.Context, logger *logrus.Entry, updateChannel chan *processing.ConfigUpdateEvent) *ManagerContext {
	return &ManagerContext{
		Context:       context,
		Logger:        logger,
		UpdateChannel: updateChannel,
	}
}
