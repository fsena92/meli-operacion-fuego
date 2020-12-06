package structs

type Configuration struct {
	Satellites []Satellite `json:"satellites"`
}

type Satellite struct {
	Name string 			`json:"name"`
	Position Position		`json:"position"`
}

/*SatelliteRequest que representa al satelite del json*/
type SatelliteRequest struct {
	Name string 			`json:"name"`
	Distance float32		`json:"distance" binding:"required"`
	Message []string 		`json:"message" binding:"required"`
}

/*Request representa la estructura del request para el post del nivel 2*/
type Request struct {
	Satellites []SatelliteRequest `json:"satellites" binding:"required,dive"`
}

/*Position representa una posicion en el plano*/
type Position struct {
	X float32       	`json:"x"`
	Y float32			`json:"y"`
}

/*Translator representa la respuesta*/
type Translator struct {
	Position Position 		`json:"position"`
	Message  string 		`json:"message"`
}

/*ResponseError representa el error que se devuelve en la api en caso de haberlo*/
type ResponseError struct {
	Description string 		`json:"description"`
}

var AllSatellites = make(map[string]Satellite)

/*SatellitesConfigured represents the list of satellites configurated in the application*/
var SatellitesConfigured []Satellite

