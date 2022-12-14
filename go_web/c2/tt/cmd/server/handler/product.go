package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_web/c2/tt/internal/products"
)

type request struct {
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
 }
 
type Product struct {
	service products.Service
}
 
func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
		   ctx.JSON(401, gin.H{
			  "error": "token inválido",
		   })
		   return
		}
   
		p, err := c.service.GetAll()
		if err != nil {
		   ctx.JSON(404, gin.H{
			  "error": err.Error(),
		   })
		   return
		}
		ctx.JSON(200, p)
	 }
   }
func (c *Product) Store() gin.HandlerFunc {
return func(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(401, gin.H{ "error": "token inválido" })
		return
	}
	var req request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
	if err != nil {
		ctx.JSON(404, gin.H{ "error": err.Error() })
		return
	}
	ctx.JSON(200, p)
	}
}


  
 