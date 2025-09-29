package service

import (
	"errors"
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
)

type PurchaseService interface {
	CreateTransaction(purchases []model.Purchasing, sub_transaction []model.Sub_Transaction, transaction model.Transaction, productIDs []int, discountcode []string, userID int) error
	RedeemCode(discountcode string, sellerID int, Total float32) (*model.Discount_code, error)
}

type purchaseServiceImpl struct {
	repo         repository.PurchaseRepository
	repo_product repository.ProductRepository
	repo_user    repository.UsersRepository
}

var (
	ErrAddress = errors.New("address blank")
	ErrSold    = errors.New("sold")
	ErrDelete  = errors.New("delete")

	ErrInvalidCode   = errors.New("invalid code")
	ErrInvalidSeller = errors.New("invalid seller")
	ErrBelowMinimum  = errors.New("less than minimum")
	ErrExceedLimit   = errors.New("exceed limit used")
)

func NewPurchaseSevice(repo repository.PurchaseRepository, repo_product repository.ProductRepository, repo_user repository.UsersRepository) PurchaseService {
	return &purchaseServiceImpl{repo, repo_product, repo_user}
}

func (s *purchaseServiceImpl) CreateTransaction(purchases []model.Purchasing, sub_transaction []model.Sub_Transaction, transaction model.Transaction, productIDs []int, discountcode []string, userID int) error {
	Users, err := s.repo_user.GetUserByID(userID)
	if err != nil {
		fmt.Println("s.purchase: User Error :", err)
		return err
	}
	if Users.Address == nil {
		fmt.Println("s.purchase: User Error :", err)
		return ErrAddress
	}

	addr := *Users.Address

	if Users.SubDistrict != nil {
		addr += ", " + *Users.SubDistrict
	}
	if Users.District != nil {
		addr += ", " + *Users.District
	}
	if Users.Province != nil {
		addr += ", " + *Users.Province
	}

	addr += ", " + *Users.Postal_code

	for _, purchase := range purchases {
		row, err := s.repo_product.GetByID(purchase.Product_ID)

		if err != nil {
			fmt.Println("s.purchase: Purchase Error :", err)
			return err
		}
		if !row.Selling {
			fmt.Println("s.purchase: Your product has already sold")
			return ErrSold
		}
		if !row.Delflag {
			fmt.Println("s.purchase: Your product not exiting")
			return ErrDelete
		} else {
			continue
		}
	}

	tranID, err := s.repo.CreateTransaction(transaction, userID, addr)
	if err != nil {
		fmt.Println("s.purchase: Transaction Error :", err)
		return err
	}
	err_sub_tran := s.repo.CreateSubTransaction(sub_transaction, tranID)
	if err_sub_tran != nil {
		fmt.Println("s.purchase: Sub_Transaction Error :", err_sub_tran)
		return err_sub_tran
	}
	err_purchase := s.repo.CreatePurchasing(purchases, tranID)
	if err_purchase != nil {
		fmt.Println("s.purchase: Purchaselist Error :", err_purchase)
		return err_purchase
	}
	err_product := s.repo.UpdateProductSelling(productIDs)
	if err_product != nil {
		fmt.Println("s.purchase: Purchaselist Error :", err_purchase)
		return err_product
	}

	err_discountcode := s.repo.UpdateUsedCode(discountcode)
	if err_discountcode != nil {
		fmt.Println("s.purchase: DiscountCode Error :", err_discountcode)
		return err_discountcode
	}
	return nil
}

func (s *purchaseServiceImpl) RedeemCode(discountcode string, sellerID int, Total float32) (*model.Discount_code, error) {
	row, err := s.repo.RedeemCode(discountcode)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve discount code: %w", err)
	}
	if row == nil {
		return nil, ErrInvalidCode
	}
	if row.Seller_ID != sellerID {
		return nil, ErrInvalidSeller
	}
	if Total < row.Minimum_total {
		return nil, ErrBelowMinimum
	}
	if row.Used >= row.Limit {
		return nil, ErrExceedLimit
	}
	return row, nil
}
