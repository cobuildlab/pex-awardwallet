package awardwallet

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Client Cliente de la API
type Client struct {
	host   string       // URL de la API
	apikey string       // API Key
	client *http.Client // Cliente HTTP para realizar consultas
}

//NewClient Constructor Client
func NewClient(host, apiKey string) *Client {
	return &Client{
		host:   host,
		apikey: apiKey,
		client: &http.Client{},
	}
}

//GetAccountDetails Obtiene la informacion de la cuenta
func (c Client) GetAccountDetails() (awardResponse *ResponseDetailAccount, awardErr *Error, err error) {
	req, err := http.NewRequest("GET", c.host, nil)
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

//GetConnectionLink ...
func (c Client) GetConnectionLink(body []byte) (awardResponse *ResponseConnectionLink, awardErr *Error, err error) {
	req, err := http.NewRequest("POST", c.host, bytes.NewReader(body))
	if err != nil {
		return
	}

	req.Header.Set("X-Authentication", c.apikey)

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
