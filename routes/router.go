package routes

import (
	"sync"
	"time"

	"github.com/cobuildlab/pex-awardwallet/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	onceRouter sync.Once
	err        error
	router     *gin.Engine
)

//SetupRouter Get router app
//	@return *gin.Engine :
//		Router app
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//CORS Config
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}

	//Middlewares
	router.Use(cors.New(config))

	//API V1
	v1 := router.Group("/api/v1")
	{
		//AwardWallet Endpoints
		award := v1.Group("/award")
		{
			award.GET("/account", handlers.GetConnectedUser)
			award.GET("/connection", handlers.GetConnectionLink)
		}
	}

	return router
}
