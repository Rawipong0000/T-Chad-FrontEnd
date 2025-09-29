package repository

import (
	"context"
	"fmt"
	"strings"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5/pgxpool" // ✅ pgx v5
)

type HistoryRepository interface {
	GetHistoryTransaction(userID int) ([]model.History_ordering, error)
	CompleteTransaction(SubTranID int) error
	RefundTransaction(SubTranID int) error
	RefundApprove(SubTranID int) error
	RefundReject(SubTranID int) error
	CancelTransaction(SubTranID int) error
	UpdateSellingStatus(subTranID int) error
}

type historyRepoImpl struct {
	db *pgxpool.Pool
}

func NewHistoryRepository(db *pgxpool.Pool) HistoryRepository {
	return &historyRepoImpl{db}
}

func (r *historyRepoImpl) GetHistoryTransaction(userID int) ([]model.History_ordering, error) {
	// ใช้ parameterized query ($1) แทน fmt.Sprintf
	const query = `
	SELECT 
    	s.sub_tran_id,
    	s.tran_id,
    	t.transaction_id,
    	s.seller_id,
    	u.user_id AS user_user_id,
    	u.name,
		t.address,
    	u.shopname,
    	p.transaction_id AS purchase_transaction_id,
    	p.product_id AS purchase_product_id,
    	pr.product_id AS product_product_id,
    	pr.product_name,
    	pr.user_id   AS product_user_id,
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
	JOIN transaction t ON s.tran_id = t.transaction_id
	JOIN "user" u      ON s.seller_id = u.user_id
	JOIN purchasing p  ON p.transaction_id = s.tran_id
	JOIN products pr   ON p.product_id = pr.product_id AND pr.user_id = s.seller_id
	JOIN status_code st ON s.status_code = st.status_code
	WHERE t.user_id = $1;
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.History_ordering
	for rows.Next() {
		var t model.History_ordering
		if err := rows.Scan(
			// ต้องเรียงให้ตรงกับ SELECT ข้างบน
			&t.Sub_Tran_id,         // s.sub_tran_id
			&t.Sub_Tran_Tran_id,    // s.tran_id
			&t.Transaction_Tran_id, // t.transaction_id
			&t.Sub_Tran_Seller_id,  // s.seller_id
			&t.User_User_ID,        // u.user_id
			&t.Name,                // u.name
			&t.Address,
			&t.Shopname,            // u.shopname
			&t.Purchase_Tran_id,    // p.transaction_id
			&t.Purchase_Product_ID, // p.product_id
			&t.Product_Product_ID,  // pr.product_id
			&t.Product_Name,        // pr.product_name
			&t.Product_User_ID,     // pr.user_id  << เพิ่มอันนี้ให้ตรงกับ SELECT
			&t.Discount_code,       // s.discount_code
			&t.Tracking,            // s.tracking
			&t.Sub_Total,           // s.sub_total
			&t.Sub_Status_code,     // s.status_code
			&t.Status_Status_code,  // st.status_code
			&t.Status_name,         // st.status_name
			&t.Color,
			&t.Create_date, // s.create_date
			&t.Update_date, // s.update_date
			&t.Delflag,     // s.delflag
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *historyRepoImpl) CompleteTransaction(SubTranID int) error {
	query := `UPDATE sub_transaction SET status_code=$1, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$2`
	_, err := r.db.Exec(context.Background(), query, 2, SubTranID)
	return err
}

func (r *historyRepoImpl) RefundTransaction(SubTranID int) error {
	query := `UPDATE sub_transaction SET status_code=$1, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$2`
	_, err := r.db.Exec(context.Background(), query, 3, SubTranID)
	return err
}

func (r *historyRepoImpl) RefundApprove(SubTranID int) error {
	query := `UPDATE sub_transaction SET status_code=$1, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$2`
	_, err := r.db.Exec(context.Background(), query, 4, SubTranID)
	return err
}

func (r *historyRepoImpl) RefundReject(SubTranID int) error {
	query := `UPDATE sub_transaction SET status_code=$1, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$2`
	_, err := r.db.Exec(context.Background(), query, 0, SubTranID)
	return err
}

func (r *historyRepoImpl) CancelTransaction(SubTranID int) error {
	query := `UPDATE sub_transaction SET status_code=$1, update_date = CURRENT_TIMESTAMP WHERE sub_tran_id=$2`
	_, err := r.db.Exec(context.Background(), query, 5, SubTranID)
	return err
}

func (r *historyRepoImpl) UpdateSellingStatus(subTranID int) error {
	ctx := context.Background()

	// 1) ดึง product_id ทั้งหมดที่อยู่ใน sub_transaction นี้
	const q = `
		SELECT DISTINCT p.product_id
		FROM sub_transaction s
		JOIN purchasing p ON p.transaction_id = s.tran_id
		WHERE s.sub_tran_id = $1;
	`

	rows, err := r.db.Query(ctx, q, subTranID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var productIDs []int
	for rows.Next() {
		var pid int
		if err := rows.Scan(&pid); err != nil {
			return err
		}
		productIDs = append(productIDs, pid)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	if len(productIDs) == 0 {
		return nil // ไม่มีสินค้าที่ต้องอัปเดต
	}

	// 2) สร้าง placeholders และอัปเดตสถานะขาย
	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	updateQ := fmt.Sprintf(
		`UPDATE products SET selling = TRUE, update_date = CURRENT_TIMESTAMP WHERE product_id IN (%s)`,
		strings.Join(placeholders, ","),
	)

	_, err = r.db.Exec(ctx, updateQ, args...)
	return err
}
