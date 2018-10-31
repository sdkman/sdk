package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPull(t *testing.T) {
	expected := "1.0.0"

	var response = fmt.Sprintf(`{"version": "%s"}`, expected)
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()

	actual, err := Pull(apiStub.URL)
	check(err)

	assert.Equal(t, expected, actual)
}

func TestPullMalformedJson(t *testing.T) {
	var response = "invalid json"
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()

	_, err := Pull(apiStub.URL)

	assert.Error(t, err)
}

func TestPullHttpGetConnectionRefused(t *testing.T) {
	actual, err := Pull("http://localhost:0")

	assert.Equal(t, "", actual)
	assert.Error(t, err, "Connection refused")
}

func TestPullHttpGetReadTimeout(t *testing.T) {
	var response = fmt.Sprint(`{"version": "too late"}`)
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
	}))
	defer apiStub.Close()

	actual, err := Pull(apiStub.URL)

	assert.Equal(t, "", actual)
	assert.Error(t, err)
}

func TestPullHttpGetNonOkStatus(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer apiStub.Close()

	actual, err := Pull(apiStub.URL)

	assert.Equal(t, "", actual)
	assert.Error(t, err)
}
