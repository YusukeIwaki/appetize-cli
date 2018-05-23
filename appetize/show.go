package appetize

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ShowOptions struct {
	PublicKey string
}

type ShowResponse struct {
	*AppItem
	AppVersionCode string `json:appVersionCode`
	AppVersionName string `json:appVersionName`
	Bundle         string `json:bundle`
	IconUrl        string `json:iconUrl`
	Name           string `json:name`
	Note           string `json:note`
	Disabled       bool   `json:disabled`
}

func (showResponse *ShowResponse) ViewUrl() string {
	return "https://appetize.io/app/" + showResponse.PublicKey
}

// ref: https://appetize.io/docs#listing-apps
func (client *Client) ShowApp(options ShowOptions) (*ShowResponse, error) {
	req, err := client.NewApiRequest("GET", "/v1/apps/"+options.PublicKey, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	var showResponse ShowResponse
	if err := json.NewDecoder(res.Body).Decode(&showResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &showResponse, nil
}
