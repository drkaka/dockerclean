package cmd

import (
	"testing"

	"github.com/drkaka/dockerclean/req"
	"github.com/stretchr/testify/assert"
)

func TestGetCreateTimeRequestWithSuccessfulResponse(t *testing.T) {
	demoResponse := `{"history":[{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.402230769Z\"}"},{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.165340448Z\"}"}]}`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	ts, err := getTagCreateTime(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, ts, int64(1515532381402230769), "returned ts wrong")
}

func TestGetCreateTimeRequestWithInvalidJSON(t *testing.T) {
	demoResponse := `{"history":[{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.402230769Z\"}"},{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.165340448Z\"}"}]`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	ts, err := getTagCreateTime(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should have error")
	assert.Equal(t, ts, int64(0), "returned ts should be 0 for invalid JSON")
}

func TestGetCreateTimeRequestWithInvalidV1Compatibility(t *testing.T) {
	demoResponse := `{"history":[{"v1Compatibility":"\"created\":\"2018-01-09T21:13:01.402230769Z\"}"},{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.165340448Z\"}"}]}`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	ts, err := getTagCreateTime(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should have error")
	assert.Equal(t, ts, int64(0), "returned ts should be 0 for invalid V1Compatibility")
}

func TestGetCreateTimeRequestWithInvalidTimeFormat(t *testing.T) {
	demoResponse := `{"history":[{"v1Compatibility":"{\"created\":\"2018-01-0921:13:01.402230769Z\"}"},{"v1Compatibility":"{\"created\":\"2018-01-09T21:13:01.165340448Z\"}"}]}`
	mockResponse := HTTPClientResponseMock{
		Message: demoResponse,
	}

	ts, err := getTagCreateTime(&mockResponse, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should have error")
	assert.Equal(t, ts, int64(0), "returned ts should be 0 for invalid TimeFormat")
}

func TestGetCreateTimeWithInvalidURL(t *testing.T) {
	badURL := "http:example.com"
	ts, err := getTagCreateTime(req.GetClient(10), badURL, "image", "tag")
	assert.Error(t, err, "should have error")
	assert.Equal(t, ts, int64(0), "returned ts should be 0 for invalid TimeFormat")
}

func TestGetCreateTimeWithBadStatusCode(t *testing.T) {
	badStatusCode := HTTPClientStatusMock{
		Code: 201,
	}
	ts, err := getTagCreateTime(&badStatusCode, "http://dockerregistry.com:1123", "image", "tag")
	assert.Error(t, err, "should return an error for bad status code")
	assert.Equal(t, ts, int64(0), "returned ts should be 0 for bad status code")
}

func TestGetCreateTimeWithResponseError(t *testing.T) {
	responseErr := HTTPClientResponseErrorMock{}
	ts, err := getTagCreateTime(&responseErr, "http://dockerregistry.com:1123", "image", "tag")
	assert.Equal(t, int64(0), ts, "time should 0")
	assert.Equal(t, errResp, err, "should return an error for error response")
}
