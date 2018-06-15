package cmd

import (
	"net/http"
	"testing"

	"github.com/drkaka/dockerclean/req"
	"github.com/stretchr/testify/assert"
)

func TestGetTagDigestRequestWithSuccessfulResponse(t *testing.T) {
	expectingDigest := "sha256:jqk"
	header := make(http.Header, 0)
	header.Add("Docker-Content-Digest", expectingDigest)

	mockResponse := HTTPClientResponseMock{
		Header: header,
	}

	digest, err := getTagDigest(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, expectingDigest, digest, "returned digest wrong")
}

func TestGetTagDigestWithInvalidURL(t *testing.T) {
	badURL := "http:example.com"
	digest, err := getTagDigest(req.GetClient(10), badURL, "image", "tag")
	assert.Error(t, err, "should return an error for bad URL")
	assert.Equal(t, "", digest, "returned digest wrong")
}

func TestGetTagDigestWithBadStatusCode(t *testing.T) {
	badStatusCode := HTTPClientStatusMock{
		Code: 201,
	}
	digest, err := getTagDigest(&badStatusCode, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should return an error for bad status code")
	assert.Equal(t, "", digest, "returned digest wrong")
}

func TestGetTagDigestWithResponseError(t *testing.T) {
	responseErr := HTTPClientResponseErrorMock{}
	digest, err := getTagDigest(&responseErr, "http://dockerregistry.com:1123", "image", "tag")
	assert.Equal(t, "", digest, "bad status code should not have repos responsed")
	assert.Equal(t, errResp, err, "should return an error for error response")
}

func TestDeleteTagRequestWithBadResponse(t *testing.T) {
	expectingDigest := "sha256:jqk"
	header := make(http.Header, 0)
	header.Add("Docker-Content-Digest", expectingDigest)

	mockResponse := HTTPClientResponseMock{
		Header: header,
	}

	err := deleteTag(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should have error because of status code 200")
}

func TestDeleteDigestRequestWithSuccessfulResponse(t *testing.T) {
	goodStatusCode := HTTPClientStatusMock{
		Code: 202,
	}

	err := deleteDigest(&goodStatusCode, "http://dockerregistry.com:1123", "image", "tag")
	assert.NoError(t, err, "should not have error")
}

func TestDeleteDigestRequestWithInvalidURL(t *testing.T) {
	badURL := "http:example.com"
	err := deleteDigest(req.GetClient(10), badURL, "image", "tag")
	assert.Error(t, err, "should return an error for bad URL")
}

func TestDeleteDigestRequestWithBadStatusCode(t *testing.T) {
	goodStatusCode := HTTPClientStatusMock{
		Code: 200,
	}

	err := deleteDigest(&goodStatusCode, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should have error when status code != 202")
}

func TestDeleteDigestRequestWithResponseError(t *testing.T) {
	responseErr := HTTPClientResponseErrorMock{}
	err := deleteDigest(&responseErr, "http://dockerregistry.com:1123", "image", "digest")
	assert.Equal(t, errResp, err, "should return an error for error response")
}
