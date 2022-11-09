package main

import (
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/cmd/server/handler"
	cnn "github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/db"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetByName())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
