package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/4geeks/pex-awardwallet/services/awardwallet"
	"github.com/gin-gonic/gin"
)

//GetAwardWalletAccount Handle to get account award wallet
func GetAwardWalletAccount(c *gin.Context) {
	clientAward := awardwallet.NewClient(os.Getenv("AWARDWALLET_HOST_ACCOUNT"), os.Getenv("AWARDWALLET_APIKEY"))

	responseAward, errAward, err := clientAward.GetAccountDetails()
	if err != nil {
		LaunchResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}
	if errAward != nil {
		LaunchResponseJSON(c, http.StatusInternalServerError, gin.H{
			"errAward": errAward,
		})

		return
	}

	LaunchResponseJSON(c, http.StatusOK, gin.H{
		"account": responseAward,
	})
}

//GetAwardWalletConnectionLink ...
func GetAwardWalletConnectionLink(c *gin.Context) {
	clientAward := awardwallet.NewClient(
		os.Getenv("AWARDWALLET_HOST_CONNECTION_LINK"),
		os.Getenv("AWARDWALLET_APIKEY"),
	)

	platform, _ := c.GetQuery("platform")
	state, _ := c.GetQuery("state")
	granularSharingValue, _ := c.GetQuery("granularSharing")
	granularSharing, _ := strconv.ParseBool(granularSharingValue)

	accessValue, _ := c.GetQuery("access")
	access, _ := strconv.Atoi(accessValue)

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

	responseAward, errAward, err := clientAward.GetConnectionLink(body)
	if err != nil {
		LaunchResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}
	if errAward != nil {
		LaunchResponseJSON(c, http.StatusInternalServerError, gin.H{
			"errAward": errAward,
		})

		return
	}

	LaunchResponseJSON(c, http.StatusOK, gin.H{
		"url": responseAward.URL,
	})
}
