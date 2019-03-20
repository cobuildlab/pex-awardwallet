package awardwallet

import (
	"encoding/json"
	"testing"
)

func TestClient(t *testing.T) {
	award := NewClient("https://business.awardwallet.com/api/export/v1/account/1", "1010")

	if award == nil {
		t.Error()
	}
}

func TestGetConnectionLink(t *testing.T) {
	clientAward := NewClient(
		"https://business.awardwallet.com/api/export/v1/create-auth-url",
		"",
	)

	bodyMap := map[string]interface{}{
		"platform": "mobile",
		"access":   1,
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		t.Error(err)
		return
	}

	responseAward, errAward, err := clientAward.GetConnectionLink(body)
	if err != nil {
		t.Error(err)
		return
	}
	if errAward != nil {
		t.Error(err)
		return
	}

	if responseAward.URL == "" {
		t.FailNow()
		return
	}
}
