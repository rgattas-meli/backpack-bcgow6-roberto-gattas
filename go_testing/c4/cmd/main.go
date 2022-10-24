package main

import (
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_testing/c4/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	err:= r.Run(":18085")
	if err != nil {
		panic(err)
	}
}
