package service

import (
	"errors"
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"

	"github.com/jackc/pgx/v5"
)

type MyShopService interface {
	GetShopNameByID(User_ID int) (*model.Users, error)
	EditShopName(shopname string, id int) error
	GetMyShopAllProducts(userID int) ([]model.Product, error)
	GetMyShopTransaction(userID int) ([]model.Myshop_ordering, error)
	EditTracking(SubTranID int, Tracking string) error
}

type myShopServiceImpl struct {
	repo repository.MyShopRepository
}

func NewMyShopService(repo repository.MyShopRepository) MyShopService {
	return &myShopServiceImpl{repo}
}

func (s *myShopServiceImpl) GetShopNameByID(User_ID int) (*model.Users, error) {
	user, err := s.repo.GetShopNameByID(User_ID)
	if err != nil {
		// กรณีไม่พบข้อมูล (เช่น No Rows)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("s.myshop: shop not found for user ID %d", User_ID)
		}

		// กรณี error อื่นๆ (เช่น connection fail, etc)
		return nil, fmt.Errorf("s.myshop: failed to get shop info: %w", err)
	}
	return user, nil
}

func (s *myShopServiceImpl) EditShopName(shopname string, id int) error {
	if shopname == "" {
		return fmt.Errorf("s.myshop: shop name cannot be empty")
	}

	err := s.repo.EditShopName(shopname, id)
	if err != nil {
		return fmt.Errorf("s.myshop: could not update shop name: %w", err)
	}

	return nil
}

func (s *myShopServiceImpl) GetMyShopAllProducts(userID int) ([]model.Product, error) {
	products, err := s.repo.GetMyShopAllProducts(userID)
	if err != nil {
		return nil, fmt.Errorf("s.myshop: fail to query products: %w", err)
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("s.myshop: No products selling: %w", err)
	}
	return products, nil
}

func (s *myShopServiceImpl) GetMyShopTransaction(userID int) ([]model.Myshop_ordering, error) {
	transactions, err := s.repo.GetMyShopTransaction(userID)
	if err != nil {
		return nil, fmt.Errorf("s.myshop: fail to query transactions: %w", err)
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("s.myshop: No transactions: %w", err)
	}
	return transactions, nil
}

func (s *myShopServiceImpl) EditTracking(SubTranID int, Tracking string) error {
	err := s.repo.EditTracking(SubTranID, Tracking)
	if err != nil {
		fmt.Println("s.myshop: SQL update error :", err)
		return err
	}
	return nil
}
