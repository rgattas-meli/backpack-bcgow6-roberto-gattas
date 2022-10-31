package products

import (
	"bytes"
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)


func createServer() *gin.Engine {
	repo:= NewRepository()
	sv := NewService(repo)
	handler := NewHandler(sv)	
	router := gin.Default()
	router.GET("/api/v1/products", handler.GetProducts)
	return router
}

func createErrorServer() *gin.Engine {
	repo:= NewMockRepository()
	sv :=NewService(repo)
	handler := NewHandler(sv)	
	router := gin.Default()
	router.GET("/api/v1/products", handler.GetProducts)
	return router
}

func createRequest_handlertest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req,_ := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func Test_Get(t *testing.T){
	//arrange
	server := createServer()

	//act

	request, response :=	createRequest_handlertest(http.MethodGet, "api/v1/products?seller_id=FEX112AC", "")
	server.ServeHTTP(response, request)

	//assert
	assert.Equal(t, response.Code,200)

	request, response =	createRequest_handlertest(http.MethodGet, "api/v1/products", "")
	server.ServeHTTP(response, request)

	//assert
	assert.Equal(t, response.Code,400)


	server2 := createErrorServer()
	request, response =	createRequest_handlertest(http.MethodGet, "api/v1/products?seller_id=foo", "")
	server2.ServeHTTP(response, request)

	//assert
	assert.Equal(t, response.Code,500)
}