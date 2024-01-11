/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package configs

type RawConfig struct {
	rawConfig []byte
}

func NewRawConfig(rawConfig []byte) *RawConfig {
	return &RawConfig{
		rawConfig: rawConfig,
	}
}

func (config *RawConfig) GetConfig() (string, error) {
	return string(config.rawConfig), nil
}
