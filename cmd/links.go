package cmd

import (
	"fmt"

	"github.com/goware/urlx"
)

const (
	catalogSubPath = "/v2/_catalog"
	tagsSubPath    = "/v2/%s/tags/list"
	tagOpSubPath   = "/v2/%s/manifests/%s"
)

// getLink with given subpath and elements
func getLink(link, subPath string, elements ...interface{}) (string, error) {
	urlLink, err := urlx.Parse(link)
	if err != nil {
		return "", err
	}

	// set the path
	if len(elements) > 0 {
		urlLink.Path = fmt.Sprintf(subPath, elements...)
	} else {
		urlLink.Path = subPath
	}

	return urlLink.String(), nil
}
