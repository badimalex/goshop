package repositories

import (
	"database/sql"
	"errors"

	"github.com/badimalex/goshop/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO products (name, price, category) VALUES ($1, $2, $3) RETURNING id",
		product.Name,
		product.Price,
		product.Category,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) SearchProducts(searchQuery string) ([]models.Product, error) {
	query := "SELECT id, name, price, category FROM products WHERE name LIKE $1"
	rows, err := r.db.Query(query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	result, err := r.db.Exec(
		"UPDATE products SET name = $1, price = $2, category = $3 WHERE id = $4",
		product.Name,
		product.Price,
		product.Category,
		product.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *ProductRepository) GetProduct(id int) (*models.Product, error) {
	var product models.Product
	err := r.db.QueryRow(
		"SELECT id, name, price, category FROM products WHERE id = $1",
		id,
	).Scan(&product.ID, &product.Name, &product.Price, &product.Category)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}
