package api

import (
	"github.com/gin-gonic/gin"
	"github.com/fsena92/meli-operacion-fuego/ship"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter inicializa las rutas
func Setup(router *gin.Engine) {
	group := router.Group("/api")
	{
	group.POST("/topsecret", ship.TopSecret)
	group.POST("/topsecret_split/:satellite_name", ship.TopSecretSplitPost)
	group.GET("/topsecret_split", ship.TopSecretSplitGet)
	}
	router.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

