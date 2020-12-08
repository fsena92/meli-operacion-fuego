package resolver

import (
	"github.com/fsena92/meli-operacion-fuego/structs"
	"errors"
	"math"
	"strings"
)

const decimals = 0.05

/*GetLocation returns the position of the ship with the satellites information */
func GetLocation(distances []float32) (float32, float32){
	solutionX, solutionY, err := TrilaterationSolution(distances)
	if err != nil {
		panic("Not found solution")
	}

	return float32(solutionX), float32(solutionY)
}
/*ValidateSatellitesRequested validates the number of the satellites, the satellites registered in the configuration*/
func ValidateSatellitesRequested(satellitesRequested []structs.SatelliteRequest) bool{
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
/*ValidateSatelliteNameExistsInConfig validates the satelite name in the request into the satellites configurated*/
func ValidateSatelliteNameExistsInConfig(satelliteName string) bool{
	return contains(structs.SatellitesConfigured, strings.ToLower(satelliteName))
}

func contains(elements []structs.Satellite, name string) bool {
	for _, element := range elements {
		if element.Name == name {
			return true
		}
	}
	return false
}

/*ValidateLocation validates the point returned on TrilaterarionSolution */
func ValidateLocation(x float32, y float32, distances []float32) bool{
		a1 := structs.SatellitesConfigured[0].Position.X
		b1 := structs.SatellitesConfigured[0].Position.Y

		difX := (round(float64(x), decimals) - float64(a1))
		difY := (round(float64(y), decimals) - float64(b1))
		if round(math.Sqrt(math.Pow(difX, 2) +  math.Pow(difY, 2)), 0.11) != round(float64(distances[0]), 0.11){
			return false
		}
		return true
}
/*TrilaterationSolution uses Cramer's rule with determinant matrices to find the solution of a system of linear equation with 2 variables */
func TrilaterationSolution(distances []float32)(float64, float64, error){
	
	var r1 float32 = distances[0]
	var r2 float32 = distances[1] 
	var r3 float32 = distances[2]

	var a1 = structs.SatellitesConfigured[0].Position.X
	var a2 = structs.SatellitesConfigured[1].Position.X
	var a3 = structs.SatellitesConfigured[2].Position.X

	var b1 = structs.SatellitesConfigured[0].Position.Y
	var b2 = structs.SatellitesConfigured[1].Position.Y
	var b3 = structs.SatellitesConfigured[2].Position.Y

	var (
		r1Sq = r1*r1
		r2Sq = r2*r2
		r3Sq = r3*r3
		a1Sq = a1*a1
		a2Sq = a2*a2
		a3Sq = a3*a3
		b1Sq = b1*b1
		b2Sq = b2*b2
		b3Sq = b3*b3
	)

	// The system of linear equation 
	// Ax + By = C
	// Dx + Ey = F

	A := (a2-a1)
	B := (b2-b1)
	C := (r1Sq-r2Sq-a1Sq+a2Sq-b1Sq+b2Sq)/2
	D := (a3-a2)
	E := (b3-b2)
	F := (r2Sq-r3Sq-a2Sq+a3Sq-b2Sq+b3Sq)/2

	d := createMatrix(A, B, D, E)
	d1 := createMatrix(C, B, F, E)
	d2 := createMatrix(A, C, D, F)

	det := determinant(d)
	det1 := determinant(d1)
	det2 := determinant(d2)

	if det != 0 {
		val1 := round(float64(det1 / det), decimals)
		val2 := round(float64(det2 / det), decimals)
		return val1, val2, nil
	}
	return 0, 0, errors.New("Not found a solution")
}

/* Create matrix for determinant */
func createMatrix(x1,y1,x2,y2 float32)([2][2]float32) {
	return [2][2]float32{{x1,y1}, {x2,y2},}
}

/* Determinant from 2x2 matrix */
func determinant(mat [2][2]float32) float32 {
	var ans float32
	ans = mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	return ans
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
