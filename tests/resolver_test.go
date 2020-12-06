package tests

import (
	"github.com/fsena92/meli-operacion-fuego/resolver"
	"github.com/fsena92/meli-operacion-fuego/structs"
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
)

func TestMain(m *testing.M){
	
	structs.SatellitesConfigured = append(structs.SatellitesConfigured, structs.Satellite{Name: "kenobi", Position: structs.Position{X: -500, Y: -200,},})
	structs.SatellitesConfigured = append(structs.SatellitesConfigured, structs.Satellite{Name: "skywalker", Position: structs.Position{X: 100, Y: -100,},})
	structs.SatellitesConfigured = append(structs.SatellitesConfigured, structs.Satellite{Name: "sato", Position: structs.Position{X: 500, Y: 100,},})
	exitVal := m.Run()
	os.Exit(exitVal)
}
func TestLocationWithPointWithEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 485.41, 265.75, 600.52)
	var x,y float32 = resolver.GetLocation(distances)
	assert.Equal(t, float32(-100.000000), x, "They should be equal")
	assert.Equal(t, float32(75.0000000), y, "They should be equal")
}

func TestLocationWithPoint2WithEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 500, 282.84, 600)
	var x,y float32 = resolver.GetLocation(distances)
	assert.Equal(t, float32(-100.0), x, "They should be equal")
	assert.Equal(t, float32(100.0), y, "They should be equal")
}

func TestLocationWithPoint3WithNotEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 100, 282.84, 600)
	var x,y float32 = resolver.GetLocation(distances)
	assert.NotEqual(t, float32(-100.0), x, "They should not be equal")
	assert.NotEqual(t, float32(100.0), y, "They should not be equal")
}

func TestLocationWithPointWithInvalidDistanceAndCoordinates(t *testing.T){
	var distances []float32
	distances = append(distances, 485.7, 266.1, 600)
	var x,y float32 = resolver.GetLocation(distances)
	assert.NotEqual(t, float32(-100.0), x)
	assert.NotEqual(t, float32(75.0), y)
}
