package service

import (
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
)

type HistoryService interface {
	GetHistoryTransaction(userID int) ([]model.History_ordering, error)
	CompleteTransaction(SubTranID int) error
	RefundTransaction(SubTranID int) error
	RefundApprove(SubTranID int) error
	RefundReject(SubTranID int) error
	CancelTransaction(SubTranID int) error
}

type historyServiceImpl struct {
	repo repository.HistoryRepository
}

func NewHistoryService(repo repository.HistoryRepository) HistoryService {
	return &historyServiceImpl{repo}
}

func (s *historyServiceImpl) GetHistoryTransaction(userID int) ([]model.History_ordering, error) {
	transactions, err := s.repo.GetHistoryTransaction(userID)
	if err != nil {
		return nil, fmt.Errorf("s.history: fail to query transactions: %w", err)
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("s.history: No transactions: %w", err)
	}
	return transactions, nil
}

func (s *historyServiceImpl) CompleteTransaction(SubTranID int) error {
	err := s.repo.CompleteTransaction(SubTranID)
	if err != nil {
		fmt.Println("s.history.CompleteTransaction: SQL update error :", err)
		return err
	}
	return nil
}

func (s *historyServiceImpl) RefundTransaction(SubTranID int) error {
	err := s.repo.RefundTransaction(SubTranID)
	if err != nil {
		fmt.Println("s.history.RefundTransaction: SQL update error :", err)
		return err
	}
	return nil
}

func (s *historyServiceImpl) RefundApprove(SubTranID int) error {
	err := s.repo.RefundApprove(SubTranID)
	if err != nil {
		fmt.Println("s.history.RefundApprove: SQL update error :", err)
		return err
	}
	err_pr_st := s.repo.UpdateSellingStatus(SubTranID)
	if err_pr_st != nil {
		fmt.Println("s.history.RefundApprove: SQL update product status error :", err)
		return err_pr_st
	}

	return nil
}

func (s *historyServiceImpl) RefundReject(SubTranID int) error {
	err := s.repo.RefundReject(SubTranID)
	if err != nil {
		fmt.Println("s.history.RefundReject: SQL update error :", err)
		return err
	}
	return nil
}

func (s *historyServiceImpl) CancelTransaction(SubTranID int) error {
	err := s.repo.CancelTransaction(SubTranID)
	if err != nil {
		fmt.Println("s.history.CancelTransaction: SQL update error :", err)
		return err
	}
	err_pr_st := s.repo.UpdateSellingStatus(SubTranID)
	if err_pr_st != nil {
		fmt.Println("s.history.RefundApprove: SQL update product status error :", err)
		return err_pr_st
	}

	return nil
}
