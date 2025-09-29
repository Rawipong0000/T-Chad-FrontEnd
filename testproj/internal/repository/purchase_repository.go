package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PurchaseRepository interface {
	CreatePurchasing(purchases []model.Purchasing, TranID int) error
	CreateTransaction(transaction model.Transaction, userID int, address string) (int, error)
	CreateSubTransaction(sub_transaction []model.Sub_Transaction, tranID int) error
	UpdateProductSelling(productIDs []int) error
	RedeemCode(discountcode string) (*model.Discount_code, error)
	UpdateUsedCode(discountCode []string) error
}

type purchaseRepoImpl struct {
	db *pgxpool.Pool
}

func NewPurchaseRepository(db *pgxpool.Pool) PurchaseRepository {
	return &purchaseRepoImpl{db}
}

func (r *purchaseRepoImpl) CreatePurchasing(purchases []model.Purchasing, tranID int) error {
	if len(purchases) == 0 {
		return nil
	}

	values := []string{}
	args := []interface{}{tranID} // $1 = tranID

	for i, purchase := range purchases {
		argIndex := i*2 + 2       // product_id
		argIndex2 := argIndex + 1 // discount_code
		values = append(values, fmt.Sprintf("($1, $%d, $%d)", argIndex, argIndex2))
		args = append(args, purchase.Product_ID, purchase.Discount_code)
	}

	query := fmt.Sprintf(`
    INSERT INTO purchasing (transaction_id, product_id, discount_code)
    VALUES %s
`, strings.Join(values, ","))

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (r *purchaseRepoImpl) CreateTransaction(transaction model.Transaction, userID int, address string) (int, error) {
	query := `INSERT INTO transaction (user_id,discount,total,address) 
	VALUES ($1, $2, $3, $4)
	RETURNING transaction_id`

	var newTran int

	err := r.db.QueryRow(context.Background(), query,
		userID, transaction.Discount, transaction.Total, address).
		Scan(&newTran)

	if err != nil {
		return 0, err
	}

	return newTran, err
}

func (r *purchaseRepoImpl) CreateSubTransaction(sub_transaction []model.Sub_Transaction, tranID int) error {
	if len(sub_transaction) == 0 {
		return nil
	}

	values := []string{}
	args := []interface{}{tranID} // $1 = tranID

	for i, sub_transaction := range sub_transaction {
		argIndex := i*3 + 2       // product_id
		argIndex2 := argIndex + 1 // discount_code
		argIndex3 := argIndex2 + 1
		values = append(values, fmt.Sprintf("($1, $%d, $%d, $%d)", argIndex, argIndex2, argIndex3))
		args = append(args, sub_transaction.Seller_ID, sub_transaction.Discount_code, sub_transaction.Sub_Total)
	}

	query := fmt.Sprintf(`
    INSERT INTO sub_transaction (tran_id, seller_id, discount_code, sub_total)
    VALUES %s
`, strings.Join(values, ","))

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (r *purchaseRepoImpl) UpdateProductSelling(productIDs []int) error {
	if len(productIDs) == 0 {
		return nil
	}

	placeholder := make([]string, len(productIDs))
	arg := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		placeholder[i] = fmt.Sprintf("$%d", i+1)
		arg[i] = id
	}

	query := fmt.Sprintf(`UPDATE products SET selling = false, update_date = CURRENT_TIMESTAMP WHERE product_id IN (%s)`, strings.Join(placeholder, ","))
	_, err := r.db.Exec(context.Background(), query, arg...)
	return err
}

func (r *purchaseRepoImpl) RedeemCode(discountcode string) (*model.Discount_code, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT discount_id,seller_id, "limit", used, discount_by_percent, discount_by_number, minimum_total, maximum_discount, delflag
		 FROM discount_code WHERE discount_code = $1`, discountcode)

	var code model.Discount_code
	err := row.Scan(&code.Discount_ID, &code.Seller_ID, &code.Limit, &code.Used, &code.Discount_by_percent, &code.Discount_by_number, &code.Minimum_total, &code.Maximum_discount,
		&code.Delflag)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &code, nil // ✅ user เจอ → คืนข้อมูล
}

func (r *purchaseRepoImpl) UpdateUsedCode(discountCode []string) error {
	if len(discountCode) == 0 {
		return nil
	}

	placeholder := make([]string, len(discountCode))
	arg := make([]interface{}, len(discountCode))
	for i, code := range discountCode {
		placeholder[i] = fmt.Sprintf("$%d", i+1)
		arg[i] = code
	}

	query := fmt.Sprintf(`UPDATE discount_code SET used = used+1 , update_date = CURRENT_TIMESTAMP WHERE discount_code IN (%s)`, strings.Join(placeholder, ","))
	_, err := r.db.Exec(context.Background(), query, arg...)
	return err
}
