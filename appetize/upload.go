package appetize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/pkg/errors"
)

type UploadOptions struct {
	AbsFilePath string
	Platform    string
}

type UploadResponse struct {
	PublicKey   string    `json:publicKey`
	Created     time.Time `json:created`
	Updated     time.Time `json:updated`
	Platform    string    `json:platform`
	VersionCode int       `json:versionCode`
}

// ref: https://appetize.io/docs#direct-uploads
func (client Client) Upload(options UploadOptions) (*UploadResponse, error) {
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

	req, err := client.NewApiRequest("POST", "/v1/apps", body)
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
	if err := writer.WriteField("platform", platform); err != nil {
		return "", err
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
