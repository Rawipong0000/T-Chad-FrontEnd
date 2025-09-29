package repository

import (
	"context"
	"testproj/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MasterRepository interface {
	GetProvince(ctx context.Context) ([]model.Province, error)
	GetDistrict(ProvinceID int, ctx context.Context) ([]model.District, error)
	GetSubDistrict(DistrictID int, ctx context.Context) ([]model.SubDistrict, error)
}

type masterRepoImpl struct {
	db *pgxpool.Pool
}

func NewMasterRepository(db *pgxpool.Pool) MasterRepository {
	return &masterRepoImpl{db}
}

func (r *masterRepoImpl) GetProvince(ctx context.Context) ([]model.Province, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT province_id,name_th,name_en FROM province`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var province []model.Province
	for rows.Next() {
		var u model.Province
		if err := rows.Scan(&u.Province_ID, &u.NameTH, &u.NameEN); err != nil {
			return nil, err
		}
		province = append(province, u)
	}
	return province, nil
}

func (r *masterRepoImpl) GetDistrict(ProvinceID int, ctx context.Context) ([]model.District, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT district_id,province_id,name_th,name_en FROM district WHERE province_id = $1`, ProvinceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var district []model.District
	for rows.Next() {
		var u model.District
		if err := rows.Scan(&u.District_ID, &u.Province_ID, &u.NameTH, &u.NameEN); err != nil {
			return nil, err
		}
		district = append(district, u)
	}
	return district, nil
}

func (r *masterRepoImpl) GetSubDistrict(DistrictID int, ctx context.Context) ([]model.SubDistrict, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT subdistrict_id,district_id,name_th,name_en,lat,long,zipcode FROM subdistrict WHERE district_id = $1`, DistrictID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subdistrict []model.SubDistrict
	for rows.Next() {
		var u model.SubDistrict
		if err := rows.Scan(&u.SubDistrict_ID, &u.District_ID, &u.NameTH, &u.NameEN, &u.Lat, &u.Long, &u.Zipcode); err != nil {
			return nil, err
		}
		subdistrict = append(subdistrict, u)
	}
	return subdistrict, nil
}
