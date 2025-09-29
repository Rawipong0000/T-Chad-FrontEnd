package service

import (
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UsersService interface {
	GetUser(User_ID int) (*model.Users, error)
	GetUserForPage(User_ID int) (*model.Users, error)
	GetUserEmail(Email, Password string) (string, error)
	CheckUserEmail(Email string) (bool, error)
	GetAllUsers() ([]model.Users, error)
	UpdateUser(user model.Users) error
	CreateUser(user model.Users) error
	DeleteUser(userid int) error
}

type usersServiceImpl struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersServiceImpl{repo}
}

func (s *usersServiceImpl) GetUser(User_ID int) (*model.Users, error) {
	return s.repo.GetUserByID(User_ID)
}

func (s *usersServiceImpl) GetUserForPage(User_ID int) (*model.Users, error) {
	user, err := s.repo.GetUserByID(User_ID)
	if err != nil {
		fmt.Println("s.GetUserForPage :", err)
		return nil, err
	}
	if user == (&model.Users{}) {
		fmt.Println("invalid UserID")
		return nil, nil
	}
	return user, nil
}

func (s *usersServiceImpl) GetUserEmail(Email, Password string) (string, error) {
	user, err := s.repo.GetUserByEmail(Email)
	if err != nil {
		fmt.Println("Error :", err)
		return "", err
	}
	if user == (&model.Users{}) {
		fmt.Println("invalid Email")
		return "", nil
	}
	if user.Password != Password {
		fmt.Println("invalid Password")
		return "", nil
	}
	token, err := generateToken(user.User_ID)
	if err != nil {
		fmt.Println("Error :", err)
		return "", err
	}
	return token, nil
}

func (s *usersServiceImpl) CheckUserEmail(email string) (bool, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		fmt.Println("Error :", err)
		return false, err
	}
	if user == nil {
		fmt.Println("Available email")
		return false, nil
	} else {
		fmt.Println("This email has already used")
		return true, nil
	}
}

func (s *usersServiceImpl) GetAllUsers() ([]model.Users, error) {
	return s.repo.GetAllUsers()
}

func (s *usersServiceImpl) UpdateUser(user model.Users) error {
	err := s.repo.UpdateUser(user)
	if err != nil {
		fmt.Println("s.users.UpdateUser: SQL error :", err)
		return err
	}
	return nil
}

func (s *usersServiceImpl) CreateUser(user model.Users) error {
	return s.repo.CreateUser(user)
}

func (s *usersServiceImpl) DeleteUser(userid int) error {
	return s.repo.DeleteUser(userid)
}

var jwtSecret = []byte("FHJKDFjksdczvfvbd45")

func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
