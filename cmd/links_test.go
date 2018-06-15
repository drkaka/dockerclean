package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinkWithBadURL(t *testing.T) {
	fullLink, err := getLink("http:/192.168.1.200:5000", catalogSubPath)
	assert.Error(t, err, "should have error with bad URL")
	assert.Equal(t, "", fullLink, "full link should be empty")
}

func TestGetLinkWithoutScheme(t *testing.T) {
	fullLink, err := getLink("192.168.1.200:5000", catalogSubPath)
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, "http://192.168.1.200:5000/v2/_catalog", fullLink, "full link wrong")
}

func TestGetLinkWithHTTPS(t *testing.T) {
	fullLink, err := getLink("https://hub.docker.com/", catalogSubPath)
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, "https://hub.docker.com/v2/_catalog", fullLink, "full link wrong")
}

func TestGetLinkWithUserPass(t *testing.T) {
	fullLink, err := getLink("//abc:def@192.168.1.200", catalogSubPath)
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, "http://abc:def@192.168.1.200/v2/_catalog", fullLink, "full link wrong")
}

func TestGetLinkWithTags(t *testing.T) {
	fullLink, err := getLink("//abc:def@192.168.1.200", tagsSubPath, "drkaka/alpine")
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, "http://abc:def@192.168.1.200/v2/drkaka/alpine/tags/list", fullLink, "full link wrong")
}

func TestGetLinkWithTagOp(t *testing.T) {
	fullLink, err := getLink("//abc:def@192.168.1.200", tagOpSubPath, "drkaka/alpine", "latest")
	assert.NoError(t, err, "should not have error")
	assert.Equal(t, "http://abc:def@192.168.1.200/v2/drkaka/alpine/manifests/latest", fullLink, "full link wrong")
}
