package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstract(t *testing.T) {
	// Arrange
	num1 := 5
	num2 := 2
	esperado := 3
	// Act
	resultado, _ := substract(num1, num2)
	// Assert
	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d ", resultado, esperado)
}

func TestDivideGood(t *testing.T) {
	// Arrange
	num1 := 25
	num2 := 5
	esperado := 5.0
	// Act
	resultado, err := divide(num1, num2)
	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado, "El numero resultado: %v, es distinto del esperado: %d ", resultado, esperado)
}

func TestDivideBad(t *testing.T) {
	// Arrange
	num1 := 25
	num2 := 0
	esperado := 0.0
	// Act
	resultado, err := divide(num1, num2)
	assert.NotNil(t, err, err.Error())
	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d ", resultado, esperado)
}
