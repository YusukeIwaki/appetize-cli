package appetize

import (
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

type DeleteOptions struct {
	PublicKey string
}

type DeleteResponse struct {
	Body string
}

// ref: https://appetize.io/docs#deleting-apps
func (client *Client) DeleteApp(options DeleteOptions) (*DeleteResponse, error) {
	req, err := client.NewApiRequest("DELETE", "/v1/apps/"+options.PublicKey, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var deleteResponse DeleteResponse
	deleteResponse.Body = strings.Trim(string(bodyBytes), " \r\n\t")
	return &deleteResponse, nil
}
