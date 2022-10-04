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
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
}

var products []request

func main() {

	r := gin.Default()
	r.GET("/ping", Ping())
	pr := r.Group("/productos")
	pr.POST("/", Store())
	r.Run()
}

func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		slice := products[len(products)-1:][0]
		req.ID = slice.ID + 1
		products = append(products, req)
		c.JSON(200, req)
	}
}

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Nombre == "" && req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "el campo Nombre y Password es requerido"})
			return
		}
		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "el campo Nombre es requerido"})
			return
		}
		if req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "el campo password es requerido"})
			return
		}

		req.ID = len(products) + 1
		products = append(products, req)
		c.JSON(http.StatusOK, req)
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
