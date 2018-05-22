package appetize

import (
	"fmt"
)

type ListOptions struct {
	Key string
}

type ListResponse struct {
	HasMore bool `json:hasMore`
}

func (client *Client) ListApps(options ListOptions) (*ListResponse, error) {
	fmt.Println("key: ", options.Key)
	var listResponse ListResponse
	return &listResponse, nil
}
