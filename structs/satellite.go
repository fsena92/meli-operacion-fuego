package structs

type Satellite struct {
	Name string 			`json:"name"`
	Position Position		`json:"position"`
}

/*SatelliteRequest que representa al satelite del json*/
type SatelliteRequest struct {
	Name string 			`json:"name"`
	Distance float32		`json:"distance"`
	Message []string 		`json:"message"`
}

/*Request representa la estructura del request para el post del nivel 2*/
type Request struct {
	Satellites []SatelliteRequest `json:"satellites"`
}

/*Position representa una posicion en el plano*/
type Position struct {
	X float32       	`json:"x"`
	Y float32			`json:"y"`
}

/*Translator representa la respuesta*/
type Translator struct {
	Position Position 		`json:"position"`
	Message []string 		`json:"message"`
}

var (
	Kenobi    = Position{kenobiX, kenobiY}
	Skywalker = Position{skywalkerX, skywalkerY}
	Sato      = Position{satoX, satoY}
)

const kenobiX = -500
const kenobiY = -200

const skywalkerX = 100
const skywalkerY = -100

const satoX = 500
const satoY = 100