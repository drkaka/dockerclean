package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/drkaka/dockerclean/req"
)

type tagsResult struct {
	Name string
	Tags []string
}

// TagsCommand the command to print all tags for the given image
func TagsCommand(link, image string, timeout int) error {
	httpClient := req.GetClient(timeout)
	repos, err := tagsRequest(httpClient, link, image, timeout)
	if err != nil {
		return err
	}

	fmt.Printf("Tags for %s:\n", image)
	for _, one := range repos {
		fmt.Printf("\t%s", one)
	}

	return nil
}

// tagsRequest send request to get all tags for the given image
func tagsRequest(httpClient req.HTTPClient, link, image string, timeout int) ([]string, error) {
	urlLink, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	// set the path
	urlLink.Path = fmt.Sprintf(tagsSubPath, image)

	// create the GET request
	tagsReq, err := http.NewRequest("GET", urlLink.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("generate request error: %+v", err)
	}

	// send the request
	resp, err := httpClient.Do(tagsReq)
	if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	} else if err != nil {
		return nil, err
	}

	// decode the catelog response
	var result tagsResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Name != image {
		return nil, errors.New("image name not match")
	}

	return result.Tags, nil
}
