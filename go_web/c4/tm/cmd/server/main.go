package main


import (
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c4/tm/cmd/server/handler"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c4/tm/internal/products"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c4/tm/pkg/store"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c4/tm/pkg/web"


)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
 
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
 }
 
