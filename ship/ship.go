package ship

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fsena92/meli-operacion-fuego/structs"
	"github.com/fsena92/meli-operacion-fuego/resolver"
	"github.com/fsena92/meli-operacion-fuego/decoder"
	"github.com/fsena92/meli-operacion-fuego/cache"
	"strings"
)

// TopSecret godoc
// @Description Recibe un json array con la información de los satélites y retorna la posicion y el mensaje del emisor
// @Accept json
// @Produce json
// @Param request body structs.Request true "Datos necesarios de satélites para localizar la posición y obtener el mensaje del emisor"
// @Success 200 {object} structs.Translator
// @Failure 400 {object} structs.ResponseError
// @Failure 403
// @Failure 404	{object} structs.ResponseError
// @Failure 500 {object} structs.ResponseError
// @Router /topsecret [post]
func TopSecret(ctx *gin.Context){
	var request structs.Request
	var distances []float32
	var messages [][]string

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error with body request": err.Error()})
		return
	}
	if !resolver.ValidateSatellitesRequested(request.Satellites){
		ctx.JSON(http.StatusBadRequest, &structs.ResponseError{ Description: "Satellites missing or repeated"})
		return
	}
	
	//Order distances and messages in the satellite requested to satellites configured
	for _, satelliteConfigured := range structs.SatellitesConfigured {
			for _, satellite := range request.Satellites {
					if strings.ToLower(satellite.Name) == strings.ToLower(satelliteConfigured.Name) {
						distances = append(distances, satellite.Distance)
						messages = append(messages, satellite.Message)
					}
				} 
		}
	
	//validate distances and messages
	var translator structs.Translator
	translator.Position.X, translator.Position.Y = resolver.GetLocation(distances)
	
	if !resolver.ValidateLocation(translator.Position.X, translator.Position.Y, distances){
		ctx.JSON(http.StatusNotFound, &structs.ResponseError{ Description: "Invalid distances with coordinates"})
		return
	}

	if !decoder.ValidateMessages(messages){
		ctx.JSON(http.StatusBadRequest, &structs.ResponseError{ Description: "Invalid message or invalid message length in any satellite"})
		return
	}

	translator.Message = decoder.GetMessage(messages)

	if !decoder.ValidateMessage(translator.Message){
		ctx.JSON(http.StatusNotFound, &structs.ResponseError{ Description: "Can't translate message"})
		return
	}
	ctx.JSON(http.StatusOK, translator)
}

// TopSecretSplitPost godoc
// @Description Recibe un json con la información del satélite enviado por parámetro y retorna el nombre del mismo si fue posible almacenarlo
// @Accept json
// @Produce json
// @Param satellite_name path string true "Nombre del satélite necesario para poder localizar la posición y obtener el mensaje del emisor" 
// @Param request body structs.SatelliteRequest true "Datos necesarios del satélite para localizar la posición y obtener el mensaje del emisor"
// @Success 200 {object} structs.SatelliteRequest
// @Failure 400 {object} structs.ResponseError
// @Failure 403
// @Failure 404	{object} structs.ResponseError
// @Failure 500	{object} structs.ResponseError
// @Router /topsecret_split/{satellite_name} [post]
func TopSecretSplitPost(ctx *gin.Context){
	satelliteName := ctx.Param("satellite_name")
	var satellite structs.SatelliteRequest
	err := ctx.BindJSON(&satellite)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error with the satellites requested": err.Error()})
		return
	}
	satellite.Name = strings.ToLower(satelliteName)

	//Validate if satellite_name exists in satellites configured 
	if resolver.ValidateSatelliteNameExistsInConfig(satellite.Name){
		inserted := cache.SetCache(satellite.Name, satellite)
		if inserted {
			ctx.JSON(http.StatusOK, satellite)
		} else {
			ctx.JSON(http.StatusInternalServerError, &structs.ResponseError{ Description: "Can't save the satellite data"})
		}
		
	} else {
			ctx.JSON(http.StatusBadRequest, &structs.ResponseError{ Description: satellite.Name + " not registered in the satellites configured"})
		}
	
}


// TopSecretSplitGet godoc
// @Description Retorna la posición y el mensaje del emisor si es posible calcularlos y si se registraron los satélites necesarios
// @Accept json
// @Produce json
// @Success 200 {object} structs.Translator
// @Failure 400 {object} structs.ResponseError
// @Failure 403
// @Failure 404 {object} structs.ResponseError
// @Failure 500 {object} structs.ResponseError
// @Router /topsecret_split [get]
func TopSecretSplitGet(ctx *gin.Context){
	if cache.CountingItems() != 3 {
		ctx.JSON(http.StatusBadRequest, &structs.ResponseError{ Description: "Missing distances and messages from satellites"})
		return
	}

	var satellitesData []structs.SatelliteRequest
	var distances []float32
	var messages [][]string
	var translator structs.Translator
	for _, satellite := range structs.SatellitesConfigured{
		satelliteData, founded := cache.GetCache(strings.ToLower(satellite.Name))
		
		if !founded {
			ctx.JSON(http.StatusNotFound, &structs.ResponseError{ Description: "Satellite not found: " + satellite.Name})
			return
		}	
		//load distances and messages
		satellitesData = append(satellitesData, satelliteData)
		distances = append(distances, satelliteData.Distance)
		messages = append(messages, satelliteData.Message)
	}
	
	//Deletes all satellites cached after loaded
	cache.FlushCache()

	translator.Position.X, translator.Position.Y = resolver.GetLocation(distances)
	if !resolver.ValidateLocation(translator.Position.X, translator.Position.Y, distances){
		ctx.JSON(http.StatusNotFound, &structs.ResponseError{ Description: "Invalid distances with coordinates"})
		return
	}

	translator.Message = decoder.GetMessage(messages)
	//validate message
	if !decoder.ValidateMessage(translator.Message){
		ctx.JSON(http.StatusNotFound, &structs.ResponseError{ Description: "Can't translate message"})
		return
	}

	ctx.JSON(http.StatusOK, translator)
}


