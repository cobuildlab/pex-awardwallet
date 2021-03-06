package awardwallet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ConnectedUser de los usuarios conectados a la cuenta
func (c Client) ConnectedUsers() (awardResponse *ResponseConnectedUser, awardErr *Error, err error) {
	req, err := http.NewRequest("GET", "https://business.awardwallet.com/api/export/v1/connectedUser", nil)
	if err != nil {
		return
	}

	req.Header.Set("X-Authentication", c.apikey)

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode <= 599 {
		err = json.Unmarshal(body, &awardErr)
		return
	}

	err = json.Unmarshal(body, &awardResponse)

	return
}
