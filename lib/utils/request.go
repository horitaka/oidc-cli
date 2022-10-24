package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Response struct {
	StatusCode int
	Body       []byte
	Error      error
}

func Get(url string) Response {
	result := Response{}

	token, _ := LoadToken()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer"+token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		result.Error = errors.Wrap(err, "Failed to call API.")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.Error = errors.Wrap(err, "Failed to read response body.")
	}

	result.StatusCode = resp.StatusCode
	result.Body = body
	return result
}
