package random_image

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type RandomImageApi struct {
	url string
}

func BuildRandomImageApi(url string) *RandomImageApi {
	return &RandomImageApi{url: url}
}

func (api *RandomImageApi) GetRandomPhoto() ([]byte, error) {
	httpResp, httpErr := http.Get(api.url)
	if httpErr != nil {
		return nil, fmt.Errorf("error requesting %s: %w", api.url, httpErr)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("received non 200 response code %d", httpResp.StatusCode)
	}
	resp, readErr := ioutil.ReadAll(httpResp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("error fetching body of %s: %w", api.url, readErr)
	}
	return resp, nil
}
