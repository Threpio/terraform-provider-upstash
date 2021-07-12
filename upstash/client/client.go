package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Upstash URL
const HostURL string = "https://api.upstash.com/v1"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Email      string
	APIKey     string
}

//NewClient returns a client object that returns against the
func NewClient(host, email, apiKey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		Email:      *email,
		APIKey:     *apiKey,
	}
	if host != nil {
		c.HostURL = *host
	}
	if (*email == "") || (*apiKey == "") {
		err := errors.New("either email or apiKey not set - Provider initialisation failed")
		return nil, err
	}
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Email, c.APIKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
