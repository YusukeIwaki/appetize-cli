package appetize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/pkg/errors"
)

type UploadOptions struct {
	AbsFilePath string
	Platform    string //for create
	PublicKey   string //for update
}

type UploadResponse struct {
	*CreatedAppItem
}

// ref: https://appetize.io/docs#direct-uploads
func (client Client) Upload(options UploadOptions) (*UploadResponse, error) {
	if options.Platform == "" && options.PublicKey == "" {
		return nil, errors.New("Specify 'platform' argument for uploading a new app, or provide PublicKey for overwriting an existing app'")
	}
	fp, err := os.Open(options.AbsFilePath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to open file: %s", options.AbsFilePath))
	}
	defer fp.Close()

	body := &bytes.Buffer{}
	contentType, err := writeMultipart(body, options.Platform, options.AbsFilePath, fp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write Multipart")
	}

	url := "/v1/apps"
	if options.PublicKey != "" {
		url += "/" + options.PublicKey
	}
	req, err := client.NewApiRequest("POST", url, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Request")
	}
	req.Header.Add("Content-Type", contentType)

	res, err := client.HandleHttpRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}
	defer res.Body.Close()

	var uploadResponse UploadResponse
	if err := json.NewDecoder(res.Body).Decode(&uploadResponse); err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &uploadResponse, nil
}

func writeMultipart(body *bytes.Buffer, platform string, absFilePath string, fp *os.File) (contentType string, err error) {
	writer := multipart.NewWriter(body)
	defer writer.Close()
	if platform != "" {
		if err := writer.WriteField("platform", platform); err != nil {
			return "", err
		}
	}
	mpart, err := writer.CreateFormFile("file", absFilePath)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(mpart, fp); err != nil {
		return "", err
	}
	return writer.FormDataContentType(), nil
}
