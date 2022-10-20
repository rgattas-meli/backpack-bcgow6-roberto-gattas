package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore interface {
	Read(data interface{}) error
	Write(data interface{}) error
	ReturnModified() bool
	ReturnRead() bool
}

type MockfileStore struct {
	FilePath string
	products []Product
	read bool
	modified bool
}

func NewStore2(pathFile string, fileName string) MockStore {
	var ps []Product
	p := Product{ID: 1, Name: "aa", Type: "aa", Count: 2, Price: 2.0}
	ps = append(ps, p)
	return &MockfileStore{fileName, ps, false, false}
}

func (fs *MockfileStore) Write(data interface{}) error {
	/*fs.modified = true
	var ps []Product
	fs.Read(&ps)

	fs.products[0] = Product{ID: 1, Name: "ff", Type: "fff", Count: 3, Price: 4.0}

	fileData, err := json.MarshalIndent(fs.products, "", "  ")
	if err != nil {
		return err
	}*/

	return nil

}

func (fs *MockfileStore) Read(data interface{}) error {
	var ps []Product
	fs.read = true
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

func (fs *MockfileStore)  ReturnModified() bool	{
	return fs.modified
}
func (fs *MockfileStore) ReturnRead() bool	{
	return fs.read
}
func TestUpdateName(t *testing.T) {
	db := NewStore2("file", "./products.json")

	repo := NewRepository(db)

	j, _ := repo.Update(1, "ff", "fff", 3, 4.0)

	p := Product{1, "ff", "fff", 3, 4.0}
	read := db.ReturnRead()
	assert.True(t, read)
	assert.Equal(t, p, j)


}
