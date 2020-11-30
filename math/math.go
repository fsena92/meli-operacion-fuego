package math

import (
	structs "github.com/fsena92/meli-operacion-fuego/structs"
	tri "github.com/co60ca/trilateration"
	"fmt"
	"math"
)


/*GetLocation retorna las coordenadas x,y del emisor del mensaje*/
func GetLocation(distances []float32) (float32, float32){
	
	params := createParameters(distances)
	resultado, err := params.SolveTrilat3()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultado)
	solutionX := float32(round(resultado[0], 0.05))
	solutionY := float32(round(resultado[1], 0.05))
	fmt.Printf("%f\n", solutionX)
	fmt.Printf("%f\n", solutionY)
	return solutionX, solutionY
}

func createPoint(x float32, y float32) (point [3]float64){
	
	point[0] = float64(x)
	point[1] = float64(y)
	point[2] = float64(0)
	return point
}

func createPoint3()(point3 []tri.Point3){
	
	var kenobiPoint = createPoint(structs.Kenobi.X, structs.Kenobi.Y)
	var skywalkerPoint = createPoint(structs.Skywalker.X, structs.Skywalker.Y)
	var satoPoint = createPoint(structs.Sato.X, structs.Sato.Y)

	point3 = append(point3, kenobiPoint, skywalkerPoint, satoPoint)
	fmt.Println(point3)
	return point3
}

func createDistances(distances []float32) (distanceTo64 []float64){
	
	for _, distance := range distances {
		distanceTo64 = append(distanceTo64, float64(distance))
	}
	return distanceTo64
}

func createParameters(distances []float32)(parameters tri.Parameters3){
	var params tri.Parameters3

	params.Loc = createPoint3()
	params.Dis = createDistances(distances)

	fmt.Println(params.Dis)
	fmt.Println(params)
	return params
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
