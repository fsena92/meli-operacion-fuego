package ship

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/fsena92/meli-operacion-fuego/structs"
	"github.com/fsena92/meli-operacion-fuego/resolver"
	"github.com/fsena92/meli-operacion-fuego/decoder"
	"github.com/fsena92/meli-operacion-fuego/cache"
	"strings"
)

// TopSecret godoc
// @Description recibe un json array con la lista de satelites y retorna la posicion y el mensaje del emisor
// @Accept json
// @Produce json
// @Param Request body Request true "Datos necesarios para localizar la posicion y obtener el mensaje del emisor"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
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
	if !validateSatellitesRequested(request.Satellites){
		ctx.JSON(http.StatusBadRequest, "Satellites repeated")
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
		
	//Llamar a funcion getLocation y getMessage
	
	//validate distances and messages
	fmt.Println("distances: ", distances)
	var translator structs.Translator
	translator.Position.X, translator.Position.Y = resolver.GetLocation(distances)
	
	if !resolver.Validate_Location(translator.Position.X, translator.Position.Y, distances){
		ctx.JSON(http.StatusNotFound, "Invalid distances with coordinates")
		return
	}

	translator.Message = decoder.GetMessage(messages)

	if !decoder.Validate_Message(translator.Message){
		ctx.JSON(http.StatusNotFound, "Invalid message")
		return
	}
	//fmt.Println(request)
	ctx.JSON(http.StatusOK, translator)
}

// TopSecretSplitPost godoc
// @Description recibe un json y retorna la posicion y el mensaje del emisor si es posible calcularla
// @Accept json
// @Produce json
// @Param satellite_name path int true "Datos necesarios para localizar la posicion y obtener el mensaje del emisor"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /topsecret_split/{satellite_name} [post]
func TopSecretSplitPost(ctx *gin.Context){
	satelliteName := ctx.Param("satellite_name")
	fmt.Println(satelliteName)
	var satellite structs.SatelliteRequest
	err := ctx.BindJSON(&satellite)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error with the satellites requested": err.Error()})
		return
	}
	satellite.Name = satelliteName
	fmt.Println(satellite)

	//Validate if satellite_name exists in satellites configured 
	if contains(structs.SatellitesConfigured, strings.ToLower(satelliteName)){
		cache.SetCache(satelliteName, satellite)
	} else {
			ctx.JSON(http.StatusBadRequest, satelliteName)
		}
		
	ctx.JSON(http.StatusOK, satelliteName)
}


// TopSecretSplitGet godoc
// @Description recibe un json y retorna la posicion y el mensaje del emisor si es posible calcularla
// @Accept json
// @Produce json
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /topsecret_split [get]
func TopSecretSplitGet(ctx *gin.Context){
	fmt.Println(cache.CountingItems())
	if cache.CountingItems() != 3 {
		ctx.JSON(http.StatusBadRequest, "Missing distances and messages from satellites")
		return
	}

	var satellitesData []structs.SatelliteRequest
	var distances []float32
	var messages [][]string
	var translator structs.Translator
	for _, satellite := range structs.SatellitesConfigured{
		satelliteData, founded := cache.GetCache(strings.ToLower(satellite.Name))
		
		if !founded {
			ctx.JSON(http.StatusNotFound, "Satellite not founded")
			return
		}	
		//cargar las distancias y los mensajes
		satellitesData = append(satellitesData, satelliteData)
		distances = append(distances, satelliteData.Distance)
		messages = append(messages, satelliteData.Message)
	}

	translator.Position.X, translator.Position.Y = resolver.GetLocation(distances)
	translator.Message = decoder.GetMessage(messages)

	ctx.JSON(http.StatusOK, translator)
}


func validateSatellitesRequested(satellitesRequested []structs.SatelliteRequest) bool{

	// validate number of satellites
	if len(satellitesRequested) != 3 {
		return false
	}
	//validate distinct satellites in request
	if ((strings.ToLower(satellitesRequested[0].Name) == strings.ToLower(satellitesRequested[1].Name)) || (strings.ToLower(satellitesRequested[1].Name) == strings.ToLower(satellitesRequested[2].Name)) || (strings.ToLower(satellitesRequested[0].Name) == strings.ToLower(satellitesRequested[2].Name))){
			return false
	}

	//validate satellites registered in config
	for _, satelliteRequested := range satellitesRequested {
		if (!contains(structs.SatellitesConfigured, strings.ToLower(satelliteRequested.Name))){
			return false
		}
	}

	return true
}


func contains(elements []structs.Satellite, name string) bool {
	for _, element := range elements {
		if element.Name == name {
			return true
		}
	}
	return false
}


