package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Remote struct {
	Version string
}

// Pull fetches a remote version from a host resource
func Pull(host string) (error, string) {
	resp, err := http.Get(host)
	check(err)

	bytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	var response Remote

	err = json.Unmarshal(bytes, &response)

	return err, response.Version
}
