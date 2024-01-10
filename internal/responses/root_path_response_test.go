/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package responses

import "testing"

func TestRootPathResponse(t *testing.T) {
	response := NewRootPathResponse("Hello, World!")

	if response.Body != "Hello, World!" {
		t.Errorf("Expected response.Body to be 'Hello, World!', but was '%s'", response.Body)
	}
}
