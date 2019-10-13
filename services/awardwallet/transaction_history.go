package awardwallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//GetTransactionHistory Obtiene la informacion de la cuenta
func (c Client) GetTransactionHistory(userID int) (awardResponse *ResponseConnectedUser, awardErr *Error, err error) {
	req, err := http.NewRequest("GET", "https://business.awardwallet.com/api/export/v1/account/"+strconv.Itoa(userID), nil)
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
	fmt.Println("AWARD WALLET TRANSACTION HISTORY:", body)

	err = json.Unmarshal(body, &awardResponse)
	return
}
