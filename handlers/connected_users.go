package handlers

import (
	"github.com/cobuildlab/pex-awardwallet/services/awardwallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

//GetConnectedUser Handle to get account award wallet
func GetConnectedUsers(c *gin.Context) {
	clientAward := awardwallet.NewClient(os.Getenv("AWARDWALLET_APIKEY"))

	responseAward, errAward, err := clientAward.ConnectedUsers()
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
		"users": responseAward,
	})
}
