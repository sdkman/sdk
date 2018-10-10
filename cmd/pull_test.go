package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPullVersion(t *testing.T) {
	expected := "1.0.0"
	var response = fmt.Sprintf(`{"version": "%s"}`, expected)
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()

	err, actual := Pull(apiStub.URL)
	check(err)

	assert.Equal(t, expected, actual)
}
