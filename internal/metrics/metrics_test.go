/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package metrics

import "testing"

func TestMetrics(t *testing.T) {
	metrics, err := NewMetrics()
	if err != nil {
		t.Errorf("Error creating metrics: %v", err)
	}

	if metrics.RootPathRequestCount == nil {
		t.Errorf("RootPathRequestCount is nil")
	}

	if metrics.RootPathRequestDurations == nil {
		t.Errorf("RootPathRequestDurations is nil")
	}

	metrics.Shutdown()
}
