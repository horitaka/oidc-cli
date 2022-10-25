package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Response struct {
	StatusCode int
	Body       []byte
	Error      error
}

func Get(urlStr string, queryMap map[string]string) Response {
	result := Response{}
	token, _ := LoadToken()

	u, _ := url.Parse(urlStr)
	q := url.Values{}
	for k, v := range queryMap {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
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
