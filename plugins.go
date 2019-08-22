package connect

import "net/http"

// Plugin represents a Kafka Connect connector plugin
type Plugin struct {
	Class   string `json:"class"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// ListPlugins retrieves a list of the installed plugins.
//
// See: https://docs.confluent.io/current/connect/references/restapi.html#get--connector-plugins-
func (c *Client) ListPlugins() ([]*Plugin, *http.Response, error) {
	path := "connector-plugins"
	var names []*Plugin

	response, err := c.get(path, &names)
	return names, response, err
}
