package products

import (

	"database/sql"
	"regexp"
	"testing"


	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/internal/domains"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	//mock.ExpectPrepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )")
	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(1, 1))
	productId := 1

	repo := NewRepo(db)
	product := domains.Product{
		ID:    productId,
		Name:  "Teclado",
		Type:  "Periferico",
		Count: 2,
		Price: 340.5,
	}

	p, err := repo.Store(product)
	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product.ID, p.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestGetAllConflict(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM products")).WillReturnError(sql.ErrConnDone)

	repo := NewRepo(db)
	//ctx := context.TODO()
	result, err := repo.GetAll()

	assert.Equal(t, sql.ErrConnDone, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepo(db)
	err = repo.Delete(id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())


}


