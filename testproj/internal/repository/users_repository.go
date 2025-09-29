package repository

import (
	"context"
	"errors"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository interface {
	GetUserByID(User_ID int) (*model.Users, error)
	GetUserByEmail(email string) (*model.Users, error)
	GetAllUsers() ([]model.Users, error)
	UpdateUser(user model.Users) error
	CreateUser(user model.Users) error
	DeleteUser(userid int) error
}

type usersRepoImpl struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) UsersRepository {
	return &usersRepoImpl{db}
}

func (r *usersRepoImpl) GetUserByID(id int) (*model.Users, error) {
	row := r.db.QueryRow(context.Background(),
		"SELECT user_id,Name,Last_name,Email,password,shopname,phone,address,subdistrict,district,province,postal_code,create_date,update_date,delflag FROM \"user\" WHERE user_id = $1", id)

	var u model.Users
	if err := row.Scan(&u.User_ID, &u.Name, &u.Lastname, &u.Email, &u.Password,
		&u.Shopname, &u.Phone, &u.Address, &u.SubDistrict, &u.District, &u.Province, &u.Postal_code,
		&u.Create_date, &u.Update_date, &u.Delflag); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *usersRepoImpl) GetUserByEmail(email string) (*model.Users, error) {
	row := r.db.QueryRow(context.Background(),
		`SELECT user_id, Name, Last_name, Email, password, shopname, create_date, update_date, delflag 
		 FROM "user" WHERE Email = $1`, email)

	var u model.Users
	err := row.Scan(&u.User_ID, &u.Name, &u.Lastname, &u.Email, &u.Password,
		&u.Shopname, &u.Create_date, &u.Update_date, &u.Delflag)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // ❗ไม่เจอ user → คืน nil
		}
		return nil, err // ❌ error อื่น เช่น database ล่ม
	}

	return &u, nil // ✅ user เจอ → คืนข้อมูล
}

func (r *usersRepoImpl) GetAllUsers() ([]model.Users, error) {
	rows, err := r.db.Query(context.Background(),
		"SELECT user_id,Name,Last_name,Email,password,shopname,create_date,update_date,delflag FROM \"user\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var u model.Users
		if err := rows.Scan(&u.User_ID, &u.Name, &u.Lastname, &u.Email, &u.Password,
			&u.Shopname, &u.Create_date, &u.Update_date, &u.Delflag); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *usersRepoImpl) UpdateUser(user model.Users) error {
	query := `UPDATE "user" SET Name=$1, Last_name=$2, phone=$3, address=$4, subdistrict=$5, district=$6, province=$7, postal_code=$8, update_date = CURRENT_TIMESTAMP WHERE user_id=$9`
	_, err := r.db.Exec(context.Background(), query, user.Name, user.Lastname, user.Phone, user.Address, user.SubDistrict, user.District, user.Province, user.Postal_code, user.User_ID)
	return err
}

func (r *usersRepoImpl) CreateUser(user model.Users) error {
	query := `INSERT INTO "user" (Name,Last_name,Email,password,shopname) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(context.Background(), query, user.Name, user.Lastname, user.Email, user.Password, user.Shopname)
	return err
}

func (r *usersRepoImpl) DeleteUser(userid int) error {
	query := `DELETE FROM "user" WHERE user_id = $1`
	_, err := r.db.Exec(context.Background(), query, userid)
	return err
}
