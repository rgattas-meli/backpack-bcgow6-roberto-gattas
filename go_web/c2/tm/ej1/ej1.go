package main

/*Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
	1. Crea un endpoint mediante POST el cual reciba la entidad.
	2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
		deberán ir guardando todas las peticiones que se vayan realizando.
	3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
	buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
	nuevo registro (sin tener una variable de último ID a nivel global).	
*/
import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
 }
var products []request
var lastID int
func main() {
    r := gin.Default()

	r.POST("/productos", func(c *gin.Context) {
		var req request
    	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
	
		}
 		})
		lastID++
		req.ID = lastID
		products = append(products, req)


	
	r.Run()
	