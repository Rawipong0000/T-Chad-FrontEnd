package repository

import (
	"context"
	"fmt"
	"strings"
	"testproj/internal/model"

	"github.com/georgysavva/scany/v2/pgxscan" // ✅ v2 สำหรับ pgx v5
	"github.com/jackc/pgx/v5/pgxpool"         // ✅ pgx v5
)

type ProductRepository interface {
	GetByID(product_id int) (*model.Product, error)
	GetMultiProductsForCart(productIDs []int) ([]model.CartProduct, error)
	GetAllProducts() ([]model.Product, error)
	Update(product model.Product) error
	Create(product model.Product) error
	CreatePageProduct(product model.Product, userID int) error
	Delete(id int) error
}

type productRepoImpl struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ProductRepository {
	return &productRepoImpl{db}
}

// ///////////////////////////////// Product//////////////////////////////////
func (r *productRepoImpl) GetByID(id int) (*model.Product, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT product_id,product_name,user_id,price,description,size,img,selling,create_date,update_date,delflag FROM products WHERE product_id = $1", id)

	var u model.Product
	if err := row.Scan(&u.Product_ID, &u.Name, &u.Product_User_ID, &u.Price, &u.Description, &u.Size, &u.Img, &u.Selling, &u.Create_date, &u.Update_date, &u.Delflag); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *productRepoImpl) GetMultiProductsForCart(productIDs []int) ([]model.CartProduct, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT 
		p.product_id,
		p.product_name,
		p.user_id AS product_user_id,
		u.user_id AS user_user_id,
		u.name,
		u.shopname,
		p.price,
		p.size,
		p.img,
		p.selling,
		p.delflag 
		FROM products p LEFT JOIN "user" u ON p.user_id = u.user_id WHERE product_id IN (%s)`, strings.Join(placeholders, ","))

	var products []model.CartProduct
	err := pgxscan.Select(context.Background(), r.db, &products, query, args...)
	return products, err
}

func (r *productRepoImpl) GetAllProducts() ([]model.Product, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT 
		p.product_id,
		p.product_name,
		p.user_id AS product_user_id,
		u.user_id AS user_user_id,
		u.name,
		u.shopname,
		p.price,
		p.description,
		p.size,
		p.img,
		p.selling,
		p.create_date,
		p.update_date,
		p.delflag 
	FROM products p LEFT JOIN "user" u ON p.user_id = u.user_id WHERE selling = true`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var u model.Product
		if err := rows.Scan(&u.Product_ID, &u.Name, &u.Product_User_ID, &u.User_User_ID, &u.U_name, &u.Shopname, &u.Price, &u.Description, &u.Size, &u.Img, &u.Selling, &u.Create_date, &u.Update_date, &u.Delflag); err != nil {
			return nil, err
		}
		products = append(products, u)
	}
	return products, nil
}

func (r *productRepoImpl) Update(product model.Product) error {
	query := `UPDATE products SET product_name=$1, price=$2, description=$3, size=$4, img=$5, update_date = CURRENT_TIMESTAMP  WHERE product_id=$6`
	_, err := r.db.Exec(context.Background(), query, product.Name, product.Price, product.Description, product.Size, product.Img, product.Product_ID)
	return err
}

func (r *productRepoImpl) Create(product model.Product) error {
	query := `INSERT INTO products (product_name, user_id , price, description, size, img) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), query, product.Name, product.Product_User_ID, product.Price, product.Description, product.Size, product.Img)
	return err
}

func (r *productRepoImpl) CreatePageProduct(product model.Product, userID int) error {
	query := `INSERT INTO products (product_name, user_id , price, description, size, img) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(context.Background(), query, product.Name, userID, product.Price, product.Description, product.Size, product.Img)
	return err
}

func (r *productRepoImpl) Delete(id int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

///////////////////////////////////////////////////////////////////////
