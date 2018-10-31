package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Remote struct {
	Version string
}

// Pull fetches a remote version from a host resource
func Pull(host string) (string, error) {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(host)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response Remote
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return "", err
	}

	return response.Version, nil
}
