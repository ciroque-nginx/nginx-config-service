/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package configs

import "testing"

func TestNewRawConfig(t *testing.T) {
	configValue := []byte("test config")

	config := NewRawConfig(configValue)
	if config == nil {
		t.Error("Expected a non-nil value")
	}
}

func TestRawConfig_GetConfig(t *testing.T) {
	configValue := []byte("test config")
	config := NewRawConfig(configValue)
	if config == nil {
		t.Error("Expected a non-nil value")
	}
	configValueReturned, err := config.GetConfig()
	if err != nil {
		t.Error("Expected a nil error")
	}

	if configValueReturned != string(configValue) {
		t.Errorf("Expected %s, got %s", configValue, configValueReturned)
	}
}
