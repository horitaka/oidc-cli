package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
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

func Post(url string, bodyMap map[string]string) Response {
	result := Response{}
	token, _ := LoadToken()

	data, _ := json.Marshal(bodyMap)
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		result.Error = err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
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

func PrintResponse(resp Response) {
	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Printf("Response:\n")
	fmt.Println(string(resp.Body))
}
