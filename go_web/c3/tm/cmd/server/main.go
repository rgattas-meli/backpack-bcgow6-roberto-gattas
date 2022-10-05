package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c3/tm/cmd/server/handler"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c3/tm/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
