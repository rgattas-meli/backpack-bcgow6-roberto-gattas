package main	

import (
	"github.com/gin-gonic/gin"
	"github.com/meli-bootcamp/cmd/server/handler"
	"github.com/meli-bootcamp/internal/products"
 )

 
func main() {
	r:= gin.Default()
	pr := r.Group("/productos")
	pr.POST("/add", Store())
	r.run()


}