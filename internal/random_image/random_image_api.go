package random_image

import (
	"errors"
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

func (api *RandomImageApi) GetRandomPhoto() (resp []byte, err error) {
	httpResp, httpErr := http.Get(api.url)
	if httpErr != nil {
		err = errors.New(fmt.Sprintf("Error requesting %s: %v", api.url, httpErr))
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("Received non 200 response code %d", httpResp.StatusCode))
	}
	resp, readErr := ioutil.ReadAll(httpResp.Body)
	if readErr != nil {
		err = errors.New(fmt.Sprintf("Error fetching body of %s: %v", api.url, readErr))
	}
	return
}
