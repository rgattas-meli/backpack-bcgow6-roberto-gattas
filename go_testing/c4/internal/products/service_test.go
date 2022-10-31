package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)


type mockRepository struct {
	db []Product
}
func NewMockRepository() Repository {
	return &mockRepository{}
}

func (m *mockRepository) GetAllBySeller(sellerID string) ([]Product, error)	{
	return nil, errors.New("error")
}



func Test_GetAll(t *testing.T){
	expectedError := errors.New("error")
	sellerID := "FFF"
	mockRepository := &mockRepository{
		db: nil,
	}
	service:= NewService(mockRepository)
	product, err := service.GetAllBySeller(sellerID)

	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, product)
}