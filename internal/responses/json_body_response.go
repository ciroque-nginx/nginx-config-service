/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package responses

type JsonBodyResponse struct {
	Body string `json:"text"`
}

func NewJsonBodyResponse(json string) JsonBodyResponse {
	return JsonBodyResponse{Body: json}
}
