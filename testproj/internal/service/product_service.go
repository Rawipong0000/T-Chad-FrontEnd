package service

import (
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
)

type ProductService interface {
	GetProduct(product_id int) (*model.Product, error)
	GetMultiProductsForCart(productIDs []int) ([]model.CartProduct, error)
	GetAllProducts() ([]model.Product, error)
	UpdateProduct(product model.Product) error
	CreateProduct(product model.Product) error
	CreatePageProduct(product model.Product, userID int) error
	DeleteProduct(id int) error
}

type productServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productServiceImpl{repo}
}

// /////////////////////////// Product///////////////////////////////////
func (s *productServiceImpl) GetProduct(product_id int) (*model.Product, error) {
	return s.repo.GetByID(product_id)
}

func (s *productServiceImpl) GetMultiProductsForCart(productIDs []int) ([]model.CartProduct, error) {
	return s.repo.GetMultiProductsForCart(productIDs)
}

func (s *productServiceImpl) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productServiceImpl) UpdateProduct(product model.Product) error {
	err := s.repo.Update(product)
	if err != nil {
		fmt.Println("s.product: SQL update error :", err)
		return err
	}
	return nil
}

func (s *productServiceImpl) CreateProduct(product model.Product) error {
	return s.repo.Create(product)
}

func (s *productServiceImpl) CreatePageProduct(product model.Product, userID int) error {
	err := s.repo.CreatePageProduct(product, userID)
	if err != nil {
		fmt.Println("s.product: SQL create error :", err)
		return err
	}
	return nil
}

func (s *productServiceImpl) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

//////////////////////////////////////////////////////////////////////
