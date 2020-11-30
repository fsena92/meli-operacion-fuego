package ship

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	structs "github.com/fsena92/meli-operacion-fuego/structs"
	math "github.com/fsena92/meli-operacion-fuego/math"
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
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error with body request": err.Error()})
		return
	}

	var distances []float32
	var messages [][]string
	for _, satellite := range request.Satellites {
		distances = append(distances, satellite.Distance)
		messages = append(messages, satellite.Message)
	}
	//Llamar a funcion getLocation y getMessage
	
	x, y := math.GetLocation(distances)
	
	var translator structs.Translator
	translator.Position.X = x
	translator.Position.Y = y

	fmt.Println(request)
	ctx.JSON(http.StatusOK, translator)
}

// TopSecretSplit godoc
// @Description recibe un json y retorna la posicion y el mensaje del emisor si es posible calcularla
// @Accept json
// @Produce json
// @Param satellite_name path int true "Datos necesarios para localizar la posicion y obtener el mensaje del emisor"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /topsecret_split/{satellite_name} [post]
func TopSecretSplit(ctx *gin.Context){
	satelliteName := ctx.Param("satellite_name")
	fmt.Println(satelliteName)
	var satellite structs.SatelliteRequest
	err := ctx.BindJSON(&satellite)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	satellite.Name = satelliteName
	fmt.Println(satellite)

	ctx.JSON(200, satelliteName)
}