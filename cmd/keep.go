package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/drkaka/dockerclean/req"
)

type tagHistory struct {
	History []struct {
		V1Compatibility string
	}
}

type compatibility struct {
	Created string
}

type tagInfo struct {
	created int64
	index   int
}

// KeepCommand to delete a given of an image.
func KeepCommand(link, image string, num, timeout int) error {
	httpClient := req.GetClient(timeout)
	tags, err := tagsRequest(httpClient, link, image)
	if err != nil {
		return err
	}

	if len(tags) <= num {
		// tags count no more than num, do nothing
		fmt.Println("Done")
		return nil
	}

	// get all tags creation time
	tagInfos := make([]tagInfo, len(tags))
	for i := 0; i < len(tags); i++ {
		ts, err := getTagCreateTime(httpClient, link, image, tags[i])
		if err != nil {
			return err
		}
		tagInfos[i] = tagInfo{
			created: ts,
			index:   i,
		}
	}

	// sort from big to small
	sort.Slice(tagInfos, func(i, j int) bool {
		return tagInfos[i].created > tagInfos[j].created
	})

	// delete from index num
	for i := num; i < len(tags); i++ {
		if err := deleteTag(httpClient, link, image, tags[tagInfos[i].index]); err != nil {
			return err
		}
	}

	fmt.Println("Please run \"docker exec -it registry /bin/registry garbage-collect  /etc/docker/registry/config.yml\" to free space.")
	return nil
}

// getTagCreateTime to get the creation time of a tag using V1 API
// First using the v1 Compatibility API to get layer info: https://github.com/docker/distribution/blob/master/docs/spec/manifest-v2-1.md#manifest-field-descriptions
// Then parse the "created" field, https://docs.docker.com/engine/api/v1.37/#operation/ImageInspect
// return unix nano and error
func getTagCreateTime(httpClient req.HTTPClient, link, image, tag string) (int64, error) {
	fullLink, err := getLink(link, tagOpSubPath, image, tag)
	if err != nil {
		return 0, err
	}

	// create the GET request
	tagReq, err := http.NewRequest("GET", fullLink, nil)
	if err != nil {
		return 0, fmt.Errorf("generate request error: %+v", err)
	}
	tagReq.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v1+json")

	// send the request
	resp, err := httpClient.Do(tagReq)
	if resp != nil && resp.StatusCode != 200 {
		return 0, fmt.Errorf("bad response code: %d", resp.StatusCode)
	} else if err != nil {
		return 0, err
	}

	// get all the compatibilities
	var compas tagHistory
	err = json.NewDecoder(resp.Body).Decode(&compas)
	if err != nil {
		return 0, err
	}

	// check all history layer records and get the latest ts as the creation time.
	layout := "2006-01-02T15:04:05.000000000Z"
	ts := int64(0)
	for _, one := range compas.History {
		var compatibility compatibility
		if err := json.Unmarshal([]byte(one.V1Compatibility), &compatibility); err != nil {
			return 0, err
		}

		t, err := time.Parse(layout, compatibility.Created)
		if err != nil {
			return 0, err
		}

		if t.UnixNano() > ts {
			ts = t.UnixNano()
		}
	}

	return ts, nil
}
