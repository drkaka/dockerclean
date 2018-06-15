package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drkaka/dockerclean/req"
)

type catalogs struct {
	Repositories []string
}

// ListCommand the command to print all images
func ListCommand(link string, timeout int) error {
	httpClient := req.GetClient(timeout)
	repos, err := listRequest(httpClient, link)
	if err != nil {
		return err
	}

	fmt.Println("repositories:")
	for _, one := range repos {
		fmt.Printf("\t%s", one)
	}

	return nil
}

func listRequest(httpClient req.HTTPClient, link string) ([]string, error) {
	fullLink, err := getLink(link, catalogSubPath)
	if err != nil {
		return nil, err
	}

	// create the GET request
	catalogReq, err := http.NewRequest("GET", fullLink, nil)
	if err != nil {
		return nil, fmt.Errorf("generate request error: %+v", err)
	}

	// send the request
	resp, err := httpClient.Do(catalogReq)
	if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	} else if err != nil {
		return nil, err
	}

	// decode the catelog response
	var catalog catalogs
	if err := json.NewDecoder(resp.Body).Decode(&catalog); err != nil {
		return nil, err
	}

	return catalog.Repositories, nil
}
