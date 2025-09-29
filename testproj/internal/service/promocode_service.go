package service

import (
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
)

type PromoCodeService interface {
	GetPromoCodeByUserID(id int) ([]model.Discount_code, error)
	GetPromoCodeByID(id int) (*model.Discount_code, error)
	CreatePromoCode(NewPromoCode model.Discount_code) error
	UpdatePromoCode(UpdatePromoCode model.Discount_code) error
	DeactivatePromoCode(DiscountID int) error
}

type promoCodeServiceImpl struct {
	repo repository.PromoCodeRepository
}

func NewPromoCodeService(repo repository.PromoCodeRepository) PromoCodeService {
	return &promoCodeServiceImpl{repo}
}

func (s *promoCodeServiceImpl) GetPromoCodeByUserID(id int) ([]model.Discount_code, error) {
	Promocodes, err := s.repo.GetPromoCodeByUserID(id)
	if err != nil {
		fmt.Println("s.promocode.GetPromoCodeByUserID: ", err)
		return nil, err
	}
	return Promocodes, nil
}

func (s *promoCodeServiceImpl) GetPromoCodeByID(id int) (*model.Discount_code, error) {
	Promocode, err := s.repo.GetPromoCodeByID(id)
	if err != nil {
		fmt.Println("s.promocode.GetPromoCodeByID: ", err)
		return nil, err
	}
	return Promocode, nil
}

func (s *promoCodeServiceImpl) CreatePromoCode(NewPromoCode model.Discount_code) error {
	err := s.repo.CreatePromoCode(NewPromoCode)
	if err != nil {
		fmt.Println("s.promocode.CreatePromoCode: ", err)
		return err
	}
	return nil
}

func (s *promoCodeServiceImpl) UpdatePromoCode(UpdatePromoCode model.Discount_code) error {
	err := s.repo.UpdatePromoCode(UpdatePromoCode)
	if err != nil {
		fmt.Println("s.promocode.UpdatePromoCode: ", err)
		return err
	}
	return nil
}

func (s *promoCodeServiceImpl) DeactivatePromoCode(DiscountID int) error {
	err := s.repo.DeactivatePromoCode(DiscountID)
	if err != nil {
		fmt.Println("s.promocode.DeactivatePromoCode: ", err)
		return err
	}
	return nil
}
