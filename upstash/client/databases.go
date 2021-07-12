package client

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetDatabases() ([]Database, error) {
	req, err := http.NewRequest("GET", "https://api.upstash.com/v1/databases", nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Email, c.APIKey)

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var databases []Database
	err = json.Unmarshal(body, &databases)
	if err != nil {
		return nil, err
	}

	return databases, nil
}
