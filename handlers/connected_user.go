package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/cobuildlab/pex-awardwallet/services/awardwallet"
	"github.com/gin-gonic/gin"
)

//GetConnectedUser Handle to get account award wallet
func GetConnectedUser(c *gin.Context) {
	userID, _ := c.GetQuery("userID")
	userIDint, _ := strconv.Atoi(userID)

	clientAward := awardwallet.NewClient(os.Getenv("AWARDWALLET_APIKEY"))

	responseAward, errAward, err := clientAward.ConnectedUser(userIDint)
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
