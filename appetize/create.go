package appetize

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/YusukeIwaki/appetize-cli/optional"
)

type CreateOptions struct {
	Url       string
	Platform  string
	LaunchUrl optional.String
	Disabled  optional.Bool
	Timeout   optional.Int
	Note      optional.String
}

type CreateResponse struct {
	*CreatedAppItem
}

func (createResponse *CreateResponse) ViewUrl() string {
	return "https://appetize.io/app/" + createResponse.PublicKey
}

// ref: https://appetize.io/docs#creating-apps
func (client *Client) CreateApp(options CreateOptions) (*CreateResponse, error) {
	if options.Platform == "" {
		return nil, errors.New("Specify 'platform' argument for creating a new app")
	}
	params := map[string]interface{}{}
	params["url"] = options.Url
	params["platform"] = options.Platform
	if options.Timeout.Present() {
		timeout := options.Timeout.Get()
		acceptableValues := []int{30, 60, 90, 120, 180, 300, 600}
		for _, acceptable := range acceptableValues {
			if timeout == acceptable {
				params["timeout"] = timeout
			}
		}
	}
	if options.Note.Present() {
		note := options.Note.Get()
		if note != "null" {
			params["note"] = note
		}
	}
	if options.LaunchUrl.Present() {
		launchUrl := options.LaunchUrl.Get()
		if launchUrl != "null" {
			params["launchUrl"] = launchUrl
		}
	}
	paramsJson, err := json.MarshalIndent(params, "", "   ")

	req, err := client.NewApiRequest("POST", "/v1/apps", bytes.NewBuffer(paramsJson))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	var createResponse CreateResponse
	if err := json.NewDecoder(res.Body).Decode(&createResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &createResponse, nil
}
