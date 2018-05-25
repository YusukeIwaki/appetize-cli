package appetize

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/YusukeIwaki/appetize-cli/optional"
)

type UpdateOptions struct {
	PublicKey  string
	UpdateForm UpdateForm
}

type UpdateForm struct {
	LaunchUrl optional.String
	Disabled  optional.Bool
	Timeout   optional.Int
	Note      optional.String
}

type UpdateResponse struct {
	*AppItem
	AppVersionCode string `json:appVersionCode`
	AppVersionName string `json:appVersionName`
	Bundle         string `json:bundle`
	IconUrl        string `json:iconUrl`
	Name           string `json:name`
	Note           string `json:note`
	LaunchUrl      string `json:launchUrl`
	Disabled       bool   `json:disabled`
	Timeout        int    `json:timeout`
}

func (updateResponse *UpdateResponse) ViewUrl() string {
	return "https://appetize.io/app/" + updateResponse.PublicKey
}

// ref: https://appetize.io/docs#updating-apps
func (client *Client) UpdateApp(options UpdateOptions) (*UpdateResponse, error) {
	updateForm := options.UpdateForm
	params := map[string]interface{}{}
	if updateForm.Disabled.Present() {
		params["disabled"] = updateForm.Disabled.Get()
	}
	if updateForm.Timeout.Present() {
		timeout := updateForm.Timeout.Get()
		acceptableValues := []int{30, 60, 90, 120, 180, 300, 600}
		for _, acceptable := range acceptableValues {
			if timeout == acceptable {
				params["timeout"] = timeout
			}
		}
	}
	if updateForm.Note.Present() {
		note := updateForm.Note.Get()
		if note == "null" {
			params["note"] = nil
		} else {
			params["note"] = note
		}
	}
	if updateForm.LaunchUrl.Present() {
		launchUrl := updateForm.LaunchUrl.Get()
		if launchUrl == "null" {
			params["launchUrl"] = nil
		} else {
			params["launchUrl"] = launchUrl
		}
	}
	paramsJson, err := json.MarshalIndent(params, "", "   ")

	req, err := client.NewApiRequest("POST", "/v1/apps/"+options.PublicKey, bytes.NewBuffer(paramsJson))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	var updateResponse UpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&updateResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &updateResponse, nil
}
