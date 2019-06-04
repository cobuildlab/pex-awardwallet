package awardwallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//ConnectedUser Obtiene la informacion de la cuenta
func (c Client) ConnectedUser(userID int) (awardResponse *ResponseConnectedUser, awardErr *Error, err error) {
	req, err := http.NewRequest("GET", "https://business.awardwallet.com/api/export/v1/connectedUser/"+strconv.Itoa(userID), nil)
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
	fmt.Println("BODY -", string(body))
	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode <= 599 {
		err = json.Unmarshal(body, &awardErr)

		fmt.Println("Problem1", err)
		return
	}
	fmt.Println("Problem2", err)

	err = json.Unmarshal(body, &awardResponse)
	fmt.Println("Problem3", err)

	return
}
