package repository

import (
	"context"
	"fmt"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5/pgxpool" // ✅ pgx v5
)

type MyShopRepository interface {
	GetShopNameByID(id int) (*model.Users, error)
	EditShopName(shopname string, id int) error
	GetMyShopAllProducts(userID int) ([]model.Product, error)
	GetMyShopTransaction(userID int) ([]model.Myshop_ordering, error)
	EditTracking(SubTranID int, Tracking string) error
}

type myShopRepoImpl struct {
	db *pgxpool.Pool
}

func NewMyShopRepository(db *pgxpool.Pool) MyShopRepository {
	return &myShopRepoImpl{db}
}

func (r *myShopRepoImpl) GetShopNameByID(id int) (*model.Users, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT Name,shopname,create_date,update_date,delflag FROM \"user\" WHERE user_id = $1", id)

	var u model.Users
	if err := row.Scan(&u.Name, &u.Shopname, &u.Create_date, &u.Update_date, &u.Delflag); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *myShopRepoImpl) EditShopName(shopname string, id int) error {
	query := `UPDATE "user" SET shopname = $1, update_date = CURRENT_TIMESTAMP WHERE user_id = $2`

	_, err := r.db.Exec(context.Background(), query, shopname, id)
	if err != nil {
		return fmt.Errorf("failed to update shop name: %w", err)
	}
	return nil
}

func (r *myShopRepoImpl) GetMyShopAllProducts(userID int) ([]model.Product, error) {
	query := fmt.Sprintf(`SELECT 
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
	FROM products p LEFT JOIN "user" u ON p.user_id = u.user_id WHERE selling = true AND p.user_id = %d`, userID)
	rows, err := r.db.Query(context.Background(), query)
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

func (r *myShopRepoImpl) GetMyShopTransaction(userID int) ([]model.Myshop_ordering, error) {
	query := fmt.Sprintf(`SELECT 
		s.sub_tran_id,
		s.tran_id,
		t.transaction_id,
		t.user_id AS tranaction_user_id,
		u.user_id AS user_user_id,
		u.Name,
		t.address,
		p.transaction_id AS purchase_transaction_id,
		pr.product_id AS product_product_id,
		pr.product_name,
		pr.user_id AS product_user_id,
		s.discount_code,
		s.tracking,
		s.sub_total,
		s.status_code AS sub_status_code,
		st.status_code,
		st.status_name,
		st.color,
		s.create_date,
		s.update_date,
		s.delflag 
	FROM sub_transaction s
	INNER JOIN transaction t ON s.tran_id = t.transaction_id
	INNER JOIN "user" u ON t.user_id = u.user_id
	INNER JOIN purchasing p ON p.transaction_id = s.tran_id
	INNER JOIN products pr ON p.product_id = pr.product_id AND s.seller_id = pr.user_id
	INNER JOIN status_code st ON s.status_code = st.status_code
	WHERE s.seller_id = %d`, userID)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Myshop_ordering
	for rows.Next() {
		var t model.Myshop_ordering
		if err := rows.Scan(
			&t.Sub_Tran_id,
			&t.Sub_Tran_Tran_id,
			&t.Transaction_Tran_id,
			&t.Transaction_User_ID,
			&t.User_User_ID,
			&t.Name,
			&t.Address,
			&t.Purchase_Tran_id,
			&t.Product_Product_ID,
			&t.Product_Name,    // << product_name มาก่อน
			&t.Product_User_ID, // << แล้วค่อย product_user_id
			&t.Discount_code,
			&t.Tracking,
			&t.Sub_Total,
			&t.Sub_Status_code,
			&t.Status_Status_code,
			&t.Status_name,
			&t.Color,
			&t.Create_date,
			&t.Update_date,
			&t.Delflag,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *myShopRepoImpl) EditTracking(SubTranID int, Tracking string) error {
	query := `UPDATE sub_transaction SET tracking=$1, status_code=$2, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$3`
	_, err := r.db.Exec(context.Background(), query, Tracking, 1, SubTranID)
	return err
}
