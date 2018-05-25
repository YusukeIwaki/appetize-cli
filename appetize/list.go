package appetize

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ListOptions struct {
}

type ListResponse struct {
	Data []DetailedAppItem `json:data`
}

// ref: https://appetize.io/docs#listing-apps
func (client *Client) ListApps(options ListOptions) (*ListResponse, error) {
	req, err := client.NewApiRequest("GET", "/v1/apps", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	var listResponse ListResponse
	if err := json.NewDecoder(res.Body).Decode(&listResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &listResponse, nil
}
