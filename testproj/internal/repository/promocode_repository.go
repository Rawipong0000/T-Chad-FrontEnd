package repository

import (
	"context"
	"fmt"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PromoCodeRepository interface {
	GetPromoCodeByUserID(id int) ([]model.Discount_code, error)
	GetPromoCodeByID(id int) (*model.Discount_code, error)
	CreatePromoCode(NewPromoCode model.Discount_code) error
	UpdatePromoCode(UpdatePromoCode model.Discount_code) error
	DeactivatePromoCode(DiscountID int) error
}

type promoCodeRepoImpl struct {
	db *pgxpool.Pool
}

func NewPromoCodeRepository(db *pgxpool.Pool) PromoCodeRepository {
	return &promoCodeRepoImpl{db}
}

func (r *promoCodeRepoImpl) GetPromoCodeByUserID(id int) ([]model.Discount_code, error) {
	query := fmt.Sprintf(
		`SELECT discount_id, discount_code, "limit", used, discount_by_percent, discount_by_number, minimum_total, maximum_discount, create_date, update_date, delflag
		 FROM discount_code WHERE seller_id = %d`, id)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var PromoCodes []model.Discount_code
	for rows.Next() {
		var u model.Discount_code
		if err := rows.Scan(&u.Discount_ID, &u.Discount_code, &u.Limit, &u.Used, &u.Discount_by_percent, &u.Discount_by_number, &u.Minimum_total, &u.Maximum_discount, &u.Create_date, &u.Update_date, &u.Delflag); err != nil {
			return nil, err
		}
		PromoCodes = append(PromoCodes, u)
	}
	return PromoCodes, nil
}

func (r *promoCodeRepoImpl) GetPromoCodeByID(id int) (*model.Discount_code, error) {
	row := r.db.QueryRow(
		context.Background(),
		`SELECT discount_code, "limit", used, discount_by_percent, discount_by_number,
                minimum_total, maximum_discount, create_date, update_date, delflag
         FROM discount_code
         WHERE discount_id = $1`,
		id,
	)

	var u model.Discount_code
	if err := row.Scan(
		&u.Discount_code,
		&u.Limit,
		&u.Used,
		&u.Discount_by_percent,
		&u.Discount_by_number,
		&u.Minimum_total,
		&u.Maximum_discount,
		&u.Create_date,
		&u.Update_date,
		&u.Delflag,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *promoCodeRepoImpl) CreatePromoCode(NewPromoCode model.Discount_code) error {
	query := `INSERT INTO discount_code (seller_id , discount_code, "limit", discount_by_percent, discount_by_number, minimum_total, maximum_discount) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(context.Background(), query, NewPromoCode.Seller_ID, NewPromoCode.Discount_code, NewPromoCode.Limit, NewPromoCode.Discount_by_percent, NewPromoCode.Discount_by_number, NewPromoCode.Minimum_total, NewPromoCode.Maximum_discount)
	return err
}

func (r *promoCodeRepoImpl) UpdatePromoCode(UpdatePromoCode model.Discount_code) error {
	query := `UPDATE discount_code SET discount_code =$1,"limit" =$2 , discount_by_percent =$3, discount_by_number =$4, minimum_total =$5, maximum_discount =$6, update_date = CURRENT_TIMESTAMP WHERE discount_id =$7`
	_, err := r.db.Exec(context.Background(), query, UpdatePromoCode.Discount_code, UpdatePromoCode.Limit, UpdatePromoCode.Discount_by_percent, UpdatePromoCode.Discount_by_number, UpdatePromoCode.Minimum_total, UpdatePromoCode.Maximum_discount, UpdatePromoCode.Discount_ID)
	return err
}

func (r *promoCodeRepoImpl) DeactivatePromoCode(DiscountID int) error {
	query := `UPDATE discount_code SET delflag =false, update_date = CURRENT_TIMESTAMP WHERE discount_id =$1`
	_, err := r.db.Exec(context.Background(), query, DiscountID)
	return err
}
