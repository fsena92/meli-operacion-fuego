package tests

import (
	math "github.com/fsena92/meli-operacion-fuego/math"
	assert "github.com/stretchr/testify/assert"
	"testing"
)


func TestLocationWithPointWithEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 485.41, 265.75, 600.52)
	var x,y float32 = math.GetLocation(distances)
	assert.Equal(t, float32(-100.000000), x, "They should be equal")
	assert.Equal(t, float32(75.0000000), y, "They should be equal")
}

func TestLocationWithPoint2WithEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 500, 282.84, 600)
	var x,y float32 = math.GetLocation(distances)
	assert.Equal(t, float32(-100.0), x, "They should be equal")
	assert.Equal(t, float32(100.0), y, "They should be equal")
}

func TestLocationWithPoint3WithNotEquals(t *testing.T){
	var distances []float32
	distances = append(distances, 100, 282.84, 600)
	var x,y float32 = math.GetLocation(distances)
	assert.NotEqual(t, float32(-100.0), x, "They should not be equal")
	assert.NotEqual(t, float32(100.0), y, "They should not be equal")
}
