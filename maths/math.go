package math

import (
	//"os"
	ship "github.com/fsena92/meli-operacion-fuego/ship"
)

/*Retorna las coordenadas x,y del emisor del mensaje*/
func GetLocation(distances ...float32) (x,y float32){
	
	var r1 float32 = distances[0]
	var r2 float32 = distances[1]
	var r3 float32 = distances[2]

	var x1 = ship.Kenobi.X
	var x2 = ship.Skywalker.X
	var x3 = ship.Sato.X

	var y1 = ship.Kenobi.Latitude
	var y2 = ship.Skywalker.Latitude
	var y3 = ship.Sato.Latitude
	
	return 1,2
}