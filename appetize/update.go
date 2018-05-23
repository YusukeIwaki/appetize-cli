package appetize

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type UpdateOptions struct {
	PublicKey  string
	UpdateForm UpdateForm
}

// 無指定なのかfalseなのかを判定するためにstringにしているが、
// もっとまともなやり方が何か有る気がする
type UpdateForm struct {
	LaunchUrl string
	Disabled  string
	Note      string
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
}

func (updateResponse *UpdateResponse) ViewUrl() string {
	return "https://appetize.io/app/" + updateResponse.PublicKey
}

// ref: https://appetize.io/docs#updating-apps
func (client *Client) UpdateApp(options UpdateOptions) (*UpdateResponse, error) {
	updateForm := options.UpdateForm
	params := map[string]interface{}{}
	if updateForm.Disabled != "" {
		if updateForm.Disabled == "true" || updateForm.Disabled == "false" {
			params["disabled"] = updateForm.Disabled
		}
	}
	if updateForm.Note != "" {
		if updateForm.Note == "null" {
			params["note"] = nil
		} else {
			params["note"] = updateForm.Note
		}
	}
	if updateForm.LaunchUrl != "" {
		if updateForm.LaunchUrl == "null" {
			params["launchUrl"] = nil
		} else {
			params["launchUrl"] = updateForm.LaunchUrl
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
