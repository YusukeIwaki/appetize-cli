package appetize

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	ApiToken string
}

func (client *Client) NewApiRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("https://%s@api.appetize.io%s", client.ApiToken, endpoint)
	return http.NewRequest(method, url, body)
}

func (client *Client) HandleHttpRequest(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	return httpClient.Do(req)
}
