package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/4geeks/pex-awardwallet/services/awardwallet"
	"github.com/gin-gonic/gin"
)

//GetConnectionLink ...
func GetConnectionLink(c *gin.Context) {
	clientAward := awardwallet.NewClient(os.Getenv("AWARDWALLET_APIKEY"))

	platform, _ := c.GetQuery("platform")
	state, _ := c.GetQuery("state")
	granularSharingValue, _ := c.GetQuery("granularSharing")
	granularSharing, _ := strconv.ParseBool(granularSharingValue)

	accessValue, _ := c.GetQuery("access")
	access, _ := strconv.Atoi(accessValue)

	responseAward, errAward, err := clientAward.ConnectionLink(platform, state, access, granularSharing)
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
