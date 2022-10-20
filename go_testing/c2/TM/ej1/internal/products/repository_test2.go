package products

import (
	"encoding/json"
	"testing"
	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/go_testing/c2/TM/pkg/store"

	"github.com/stretchr/testify/assert"
)
type MockStore interface {
	Read(data interface{}) error
	Write(data interface{}) error
 }



type MockfileStore struct {
	FilePath string
	products []Product
 }

func NewStore2(pathFile string, fileName string) MockStore {
	var ps []Product
	p := Product{ID: 1, Name: "aa", Type: "aa", Count: 2,Price: 2.0}
	ps = append(ps,p)
	return &MockfileStore{fileName, ps}
 }

func (fs *MockfileStore) Write(data interface{}) error {
	return nil
}

func (fs *MockfileStore) Read(data interface{}) error {
	var ps []Product
	p := fs.products[0]
	pss := Product{2, "fffff", "aafaff", 4, 14.0}
	ps = append(ps, p)
	ps = append(ps, pss)
	fileData, err := json.MarshalIndent(ps, "", "  ")
    if err != nil {
        return err
    }
	
    return json.Unmarshal(fileData, &data)
 }

 func TestUpdateName(t *testing.T)	{
	db := NewStore2("file", "./products.json")

	repo := NewRepository(db)

	j,_:= repo.Update(1,"ff", "fff", 3,4.0)

	p := Product{1, "ff", "fff", 3, 4.0}

	assert.Equal(t,p , j)

 }