package storage

import (
	"api-products/models"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "products.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func Create(product models.Product) (models.Product, error) {
	query := `
		INSERT INTO products (name, price)
		VALUES (?, ?) 
	`

	result, err := DB.Exec(
		query,
		product.Name,
		product.Price,
	)

	if err != nil {
		return models.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}

	product.ID = int(id)

	return product, nil
}

func GetAll() ([]models.Product, error) {
	query := `
		SELECT id, name, price
		FROM products
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetById(id int) (models.Product, error) {
	query := `
		SELECT id, name, price
		FROM products
		WHERE id = ?
	`

	var product models.Product

	err := DB.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func Update(id int, product models.Product) (models.Product, bool, error) {
	query := `
		UPDATE products
		SET name = ?, price = ?
		WHERE id = ?
	`

	result, err := DB.Exec(
		query,
		product.Name,
		product.Price,
		id,
	)

	if err != nil {
		return models.Product{}, false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, false, err
	}

	if rowsAffected == 0 {
		return models.Product{}, false, nil
	}

	product.ID = id

	return product, true, nil
}

func Delete(id int) (bool, error) {
	query := `
		DELETE FROM products
		WHERE id = ?
	`

	result, err := DB.Exec(
		query,
		id,
	)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
