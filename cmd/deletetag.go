package cmd

import (
	"fmt"
	"net/http"

	"github.com/drkaka/dockerclean/req"
)

// DeleteTagCommand to delete a given of an image.
func DeleteTagCommand(link, image, tag string, timeout int) error {
	httpClient := req.GetClient(timeout)
	if err := deleteTag(httpClient, link, image, tag); err != nil {
		return err
	}
	fmt.Println("Please run \"docker exec -it registry /bin/registry garbage-collect  /etc/docker/registry/config.yml\" to free space.")
	return nil
}

// deleteTag to delete from a tag
func deleteTag(httpClient req.HTTPClient, link, image, tag string) error {
	digest, err := getTagDigest(httpClient, link, image, tag)
	if err != nil {
		return err
	}
	return deleteDigest(httpClient, link, image, digest)
}

// getTagDigest to get the digest of a tag using v2 API
func getTagDigest(httpClient req.HTTPClient, link, image, tag string) (string, error) {
	fullLink, err := getLink(link, tagOpSubPath, image, tag)
	if err != nil {
		return "", err
	}

	// create the HEAD request
	tagReq, err := http.NewRequest("HEAD", fullLink, nil)
	if err != nil {
		return "", fmt.Errorf("generate request error: %+v", err)
	}
	tagReq.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	// send the request
	resp, err := httpClient.Do(tagReq)
	if resp != nil && resp.StatusCode != 200 {
		return "", fmt.Errorf("bad response code: %d", resp.StatusCode)
	} else if err != nil {
		return "", err
	}

	return resp.Header.Get("Docker-Content-Digest"), nil
}

// deleteDigest to send a delete request with the given digest
func deleteDigest(httpClient req.HTTPClient, link, image, digest string) error {
	fullLink, err := getLink(link, tagOpSubPath, image, digest)
	if err != nil {
		return err
	}

	// create the DELETE request
	tagReq, err := http.NewRequest("DELETE", fullLink, nil)
	if err != nil {
		return fmt.Errorf("generate request error: %+v", err)
	}

	// send the request
	resp, err := httpClient.Do(tagReq)
	if resp != nil && resp.StatusCode != 202 {
		return fmt.Errorf("bad response code: %d", resp.StatusCode)
	}
	return err
}
