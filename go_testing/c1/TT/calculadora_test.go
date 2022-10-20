package calculadora

import "testing"
import "github.com/stretchr/testify/assert"

func TestSumar(t *testing.T) {
    // Se inicializan los datos a usar en el test (input/output)
    num1 := 3
    num2 := 5
    resultadoEsperado := 2

    // Se ejecuta el test
    resultado := Restar(num2, num1)

    // Se validan los resultados
    assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}