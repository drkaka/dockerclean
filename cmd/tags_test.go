package cmd

import (
	"fmt"
	"testing"

	"github.com/drkaka/dockerclean/req"
	"github.com/stretchr/testify/assert"
)

func TestTagsRequestWithSuccessfulResponse(t *testing.T) {
	image := "abcd/dcba"
	demoResponse := `{"name":"%s", "tags": ["1.2","latest"]}`
	mockResponse := HTTPClientResponseMock{
		Message: fmt.Sprintf(demoResponse, image),
	}

	tags, err := tagsRequest(&mockResponse, "http://dockerregistry.com:1123", image)
	assert.NoError(t, err, "Should not have error when getting a good response")
	assert.Equal(t, []string{"1.2", "latest"}, tags, "Tags result wrong")
}

func TestTagsRequestWithImageNotMatchResponse(t *testing.T) {
	image := "abcd/dcba"
	demoResponse := `{"name":"%s", "tags": ["1.2","latest"]}`
	mockResponse := HTTPClientResponseMock{
		Message: fmt.Sprintf(demoResponse, image),
	}

	image1 := "another name"
	tags, err := tagsRequest(&mockResponse, "http://dockerregistry.com:1123", image1)
	assert.Nil(t, tags, "bad response should not have tags responsed")
	assert.Error(t, err, "should return an error for bad response")
}

func TestTagsRequestWithInvalidJSONResponse(t *testing.T) {
	image := "abcd/dcba"
	demoResponse := `{"name":"%s", "tags": ["1.2","latest"]`
	mockResponse := HTTPClientResponseMock{
		Message: fmt.Sprintf(demoResponse, image),
	}

	tags, err := tagsRequest(&mockResponse, "http://dockerregistry.com:1123", image)
	assert.Nil(t, tags, "bad JSON should not have repos responsed")
	assert.Error(t, err, "should return an error for bad JSON")
}

func TestTagsRequestWithBadStatusCode(t *testing.T) {
	image := "abcd/dcba"
	badStatusCode := HTTPClientStatusMock{
		Code: 201,
	}
	tags, err := tagsRequest(&badStatusCode, "http://dockerregistry.com:1123", image)
	assert.Nil(t, tags, "bad status code should not have repos responsed")
	assert.Error(t, err, "should return an error for bad status code")
}

func TestTagsRequestWithInvalidURL(t *testing.T) {
	image := "abcd/dcba"
	badURL := "http:example.com"
	repos, err := tagsRequest(req.GetClient(10), badURL, image)
	assert.Nil(t, repos, "bad URL should not have repos responsed")
	assert.Error(t, err, "should return an error for bad URL")
}

func TestTagsRequestWithResponseError(t *testing.T) {
	responseErr := HTTPClientResponseErrorMock{}
	repos, err := tagsRequest(&responseErr, "http://dockerregistry.com:1123", "drkaka/alpine")
	assert.Nil(t, repos, "bad status code should not have repos responsed")
	assert.Equal(t, errResp, err, "should return an error for error response")
}
