package awardwallet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ConnectionLink ...
func (c Client) ConnectionLink(platform, state string, access int, granularSharing bool) (awardResponse *ResponseConnectionLink, awardErr *Error, err error) {
	bodyMap := map[string]interface{}{
		"platform":        platform,
		"access":          access,
		"state":           state,
		"granularSharing": granularSharing,
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		return
	}

	req, err := http.NewRequest(
		"POST",
		"https://business.awardwallet.com/api/export/v1/create-auth-url",
		bytes.NewReader(body),
	)

	if err != nil {
		return
	}

	req.Header.Set("X-Authentication", c.apikey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode <= 599 {
		err = json.Unmarshal(body, &awardErr)

		return
	}

	err = json.Unmarshal(body, &awardResponse)

	return
}
