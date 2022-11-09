package products

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rgattas-meli/backpack-bcgow6-roberto-gattas/tree/main/go_storage/c2/internal/domains"
)
var ErrNotImplemented = fmt.Errorf("not implemented")

type Repository interface {
	Store(domains.Product) (domains.Product, error)

	GetAll() ([]domains.Product, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(product domains.Product) (domains.Product, error) { // se inicializa la base

	stmt, err := r.db.Prepare("INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )") // se prepara la sentencia SQL a ejecutar
	if err != nil {
		return domains.Product{}, err
	}
	defer stmt.Close() // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price) // retorna un sql.Result y un error
	if err != nil {
		return domains.Product{}, err
	}
	insertedId, _ := result.LastInsertId() // del sql.Resul devuelto en la ejecucion obtenemos el Id insertado
	product.ID = int(insertedId)
	return product, nil
}




const (
	GetAllProducts = "SELECT * FROM products"
)

func (r *repository) GetAll() ([]domains.Product, error) {
	var products []domains.Product
	rows, err := r.db.Query(GetAllProducts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// se recorren todas las filas
	for rows.Next() {
		// por cada fila se obtiene un objeto del tipo Product
		var product domains.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Fatal(err)
			return nil, err
		}
		//se añade el objeto obtenido al slice products
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Delete(id int) error { // se inicializa la base
	stmt, err := r.db.Prepare("DELETE FROM products WHERE id=?") // se prepara la sentencia SQL a ejecutar

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()     // se cierra la sentencia al terminar. Si quedan abiertas se genera consumos de memoria
	_, err = stmt.Exec(id) // retorna un sql.Result y un error
	if err != nil {
		return err
	}

	return nil
}