package main

import (
	"github.com/fsena92/meli-operacion-fuego/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/fsena92/meli-operacion-fuego/config"
	_ "github.com/fsena92/meli-operacion-fuego/docs"
)

// @title Fire Operation Api
// @version 1.0
// @description API Restful
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host fire-operation-api.herokuapp.com
// @BasePath /api


func main() {
	config.LoadSatellites()

	router := gin.Default()
	router.Use(cors.Default())
	api.Setup(router)
	router.Run()
	

}

