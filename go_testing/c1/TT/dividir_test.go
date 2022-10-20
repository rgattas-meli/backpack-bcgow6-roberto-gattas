package calculadora

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func TestDividir(t *testing.T) {
    // Se inicializan los datos a usar en el test (input/output)
    num1 := 3
    num2 := 0
    errorEsperado := fmt.Sprintf("el denominador no puede ser %d", num2)
    // Se ejecuta el test
    resultado, err  := Dividir(num1, num2)

    // Se validan los resultados aprovechando testify
    assert.NotNil(t, resultado)
	assert.ErrorContains(t, err, errorEsperado)
}
