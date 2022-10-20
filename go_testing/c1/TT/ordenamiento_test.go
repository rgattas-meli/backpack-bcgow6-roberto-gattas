package calculadora
import "testing"
import "github.com/stretchr/testify/assert"

func TestOrd(t *testing.T) {
	a := []int{5, 3, 4, 7, 8, 9}
	b := ordenar(a)
	c := []int{3, 4, 5,7,8,	 9}
	assert.Equal(t,c,b,"deben ser iguales")

}
