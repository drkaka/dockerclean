package cmd

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drkaka/dockerclean/req"
	"github.com/stretchr/testify/assert"
)

var (
	errResp = errors.New("response error")
)

type HTTPClientStatusMock struct {
	Code int
}

func (c *HTTPClientStatusMock) Do(req *http.Request) (*http.Response, error) {
	mockResponse := http.Response{
		StatusCode: c.Code,
	}
	return &mockResponse, nil
}

type HTTPClientResponseMock struct {
	Message string
	Header  http.Header
}

func (c *HTTPClientResponseMock) Do(req *http.Request) (*http.Response, error) {
	buf := bytes.Buffer{}
	buf.WriteString(c.Message)
	mockResponse := httptest.ResponseRecorder{
		Code:      200,
		Body:      &buf,
		HeaderMap: c.Header,
	}
	return mockResponse.Result(), nil
}

type HTTPClientResponseErrorMock struct {
}

func (c *HTTPClientResponseErrorMock) Do(req *http.Request) (*http.Response, error) {
	return nil, errResp
}

func TestListRequestWithSuccessfulResponse(t *testing.T) {
	demoResponse := `{"repositories": ["a","b"]}`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	repos, err := listRequest(&mockResponse, "http://dockerregistry.com:1123")
	assert.NoError(t, err, "Should not have error when getting a good response")
	assert.Equal(t, []string{"a", "b"}, repos, "Repos result wrong")
}

func TestListRequestWithInvalidJSONResponse(t *testing.T) {
	demoResponse := `{"repositories": ["a","b"]`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	repos, err := listRequest(&mockResponse, "http://dockerregistry.com:1123")
	assert.Nil(t, repos, "bad JSON should not have repos responsed")
	assert.Error(t, err, "should return an error for bad JSON")
}

func TestListRequestWithInvalidURL(t *testing.T) {
	badURL := "http:example.com"
	repos, err := listRequest(req.GetClient(10), badURL)
	assert.Nil(t, repos, "bad URL should not have repos responsed")
	assert.Error(t, err, "should return an error for bad URL")
}

func TestListRequestWithBadStatusCode(t *testing.T) {
	badStatusCode := HTTPClientStatusMock{
		Code: 201,
	}
	repos, err := listRequest(&badStatusCode, "http://dockerregistry.com:1123")
	assert.Nil(t, repos, "bad status code should not have repos responsed")
	assert.Error(t, err, "should return an error for bad status code")
}

func TestListRequestWithResponseError(t *testing.T) {
	responseErr := HTTPClientResponseErrorMock{}
	repos, err := listRequest(&responseErr, "http://dockerregistry.com:1123")
	assert.Nil(t, repos, "bad status code should not have repos responsed")
	assert.Equal(t, errResp, err, "should return an error for error response")
}
