// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/globocom/tsuru/auth"
	"net/http"
)

// schema represents a json schema.
type schema struct {
	Title      string              `json:"title"`
	Properties map[string]property `json:"properties"`
	Required   []string            `json:"required"`
	Links      []link              `json:"links"`
}

// link represents a json schema link.
type link map[string]string

// property represents a json schema property.
type property map[string]interface{}

// appSchema returns a json schema for app.
func appSchema(w http.ResponseWriter, r *http.Request, t *auth.Token) error {
	l := []link{
		{"href": "/apps/{name}/log", "method": "GET", "rel": "log"},
		{"href": "/apps/{name}/env", "method": "GET", "rel": "get_env"},
		{"href": "/apps/{name}/env", "method": "POST", "rel": "set_env"},
		{"href": "/apps/{name}/env", "method": "DELETE", "rel": "unset_env"},
		{"href": "/apps/{name}/restart", "method": "GET", "rel": "restart"},
		{"href": "/apps/{name}", "method": "POST", "rel": "update"},
		{"href": "/apps/{name}", "method": "DELETE", "rel": "delete"},
		{"href": "/apps/{name}/run", "method": "POST", "rel": "run"},
	}
	s := schema{
		Title:    "app schema",
		Links:    l,
		Required: []string{"platform", "name"},
		Properties: map[string]property{
			"name": {
				"type": "string",
			},
			"platform": {
				"type": "string",
			},
			"ip": {
				"type": "string",
			},
			"cname": {
				"type": "string",
			},
		},
	}
	return json.NewEncoder(w).Encode(s)
}
