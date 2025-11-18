package repo

import (
	"database/sql"
	"ecommerece/domain"
	"ecommerece/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor or constructor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products (
		title,
		description,
		price,
		img_url
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

	query := `
	  SELECT 
	    id, 
	    title, 
	    description, 
	    price, 
	    img_url 
	  FROM products
	  WHERE id=$1
	`

	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1

	var prdList []*domain.Product

	query := `
	  SELECT 
	    id, 
	    title, 
	    description, 
	    price, 
	    img_url 
	  FROM products
	  LIMIT $1
	  OFFSET $2;
	`

	err := r.db.Select(&prdList, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Count() (int64, error) {
	query := `
	  SELECT COUNT(*)
	  FROM products
	`
	var count int64
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count products: %w", err)
	}
	return count, nil
}

// func (r *productRepo) Update(p Product) (*Product, error) {
// 	query := `
// 	  UPDATE products
// 	  SET title=$1, description=$2, price=$3, img_url=$4
// 	  WHERE id=$5
// 	`

// 	row := r.db.QueryRow(query, p.ID, p.Title, p.Description, p.Price, p.ImgUrl)
// 	err := row.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &p, nil
// }

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	// Use Exec instead of QueryRow for UPDATE without RETURNING
	query := `
	  UPDATE products
	  SET title=$1, description=$2, price=$3, img_url=$4
	  WHERE id=$5
	`

	result, err := r.db.Exec(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID) // updated
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected() // optional: check if row exists
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no product found with id %d", p.ID)
	}

	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
	  DELETE FROM products WHERE id =$1
	`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
