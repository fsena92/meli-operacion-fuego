package ship

/*Satellite que representa al satelite del json*/
type Satellite struct {
	Name string 			`json:"name"`
	Distance float32		`json:"distance"`
	Message []string 		`json:"message"`
}

/*Request representa la estructura del request para el post del nivel 2*/
type Request struct {
	Satellites []Satellite `json:"satellites"`
}

/*Position representa una posicion en el plano*/
type Position struct {
	X float32
	Y float32
}

/*Translator representa la respuesta*/
type Translator struct {
	Postiion Position 		`json:"position"`
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