package products

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)
type StubStore interface {
	Read(data interface{}) error
	Write(data interface{}) error
 }

 type Type string 
const (
   FileType Type = "file"
   MonoType Type = "mongo"

)


type StubfileStore struct {
	FilePath string
 }

func NewStore(pathFile string, fileName string) StubStore {
	return &StubfileStore{fileName}
 }
 func New(Stubstore Type, fileName string) StubStore {
	switch Stubstore {
	case FileType:
		return &StubfileStore{fileName}
	}
	return nil
 }

func (fs *StubfileStore) Write(data interface{}) error {
	return nil
}

func (fs *StubfileStore) Read(data interface{}) error {
	var ps []Product
	p := Product{2, "rob", "aa", 2, 11.0}
	pss := Product{1, "a", "aafaff", 4, 14.0}
	ps = append(ps, p)
	ps = append(ps, pss)
	fileData, err := json.MarshalIndent(ps, "", "  ")
    if err != nil {
        return err
    }
	
    return json.Unmarshal(fileData, &data)
 }

func TestGetAll(t *testing.T)	{
	db := NewStore("file", "./products.json")

	repo := NewRepository(db)
	j,_:= repo.GetAll()

	p := Product{2, "rob", "aa", 2, 11.0}
	pss := Product{1, "a", "aafaff", 4, 14.0}
	ps = append(ps, p)
	ps = append(ps, pss)

	assert.Equal(t,ps , j )
}