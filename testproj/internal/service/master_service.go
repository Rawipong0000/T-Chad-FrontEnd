package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testproj/internal/model"
	"testproj/internal/repository"
	"testproj/redisclient"
	"time"

	"github.com/redis/go-redis/v9"
)

type MasterService interface {
	GetProvince(ctx context.Context) ([]model.Province, error)
	GetDistrict(ProvinceID int, ctx context.Context) ([]model.District, error)
	GetSubDistrict(DistrictID int, ctx context.Context) ([]model.SubDistrict, error)
}

type masterServiceImpl struct {
	repo repository.MasterRepository
	rc   *redisclient.Client
}

func NewMasterService(repo repository.MasterRepository, rc *redisclient.Client) MasterService {
	return &masterServiceImpl{repo, rc}
}

func (s *masterServiceImpl) GetProvince(ctx context.Context) ([]model.Province, error) {
	const key = "province_redis"

	provice_redis_data, err := s.rc.Get(ctx, key)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("redis GET %q: %w", key, err)
	}

	if len(provice_redis_data) > 0 {
		var provinces []model.Province
		if u_err := json.Unmarshal([]byte(provice_redis_data), &provinces); u_err == nil {
			fmt.Println("Get Province form redis")
			return provinces, nil
		}
	}

	provinces, err := s.repo.GetProvince(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.GetProvince: %w", err)
	}

	if data, m_err := json.Marshal(provinces); m_err == nil {
		if s_err := s.rc.Set(ctx, key, data, time.Hour); s_err != nil {
			return nil, fmt.Errorf("redis SET %q: %w", key, err)
		}
	}

	return provinces, nil
}

func (s *masterServiceImpl) GetDistrict(ProvinceID int, ctx context.Context) ([]model.District, error) {
	key := fmt.Sprintf("district:province:%d", ProvinceID)

	district_redis_data, err := s.rc.Get(ctx, key)
	fmt.Println("redis data: ", district_redis_data)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("redis GET %q: %w", key, err)
	}

	if len(district_redis_data) > 0 {
		var cached []model.District
		if u_err := json.Unmarshal([]byte(district_redis_data), &cached); u_err == nil {
			fmt.Println("return cached data")
			return cached, nil
		}
	}

	district, dberr := s.repo.GetDistrict(ProvinceID, ctx)
	if dberr != nil {
		return nil, fmt.Errorf("GetDistrict repo: %w", dberr)
	}
	if len(district) == 0 {
		fmt.Println("s.users.GetDistrict: No province ID")
		return nil, nil
	}

	if data, merr := json.Marshal(district); merr == nil {
		if setErr := s.rc.Set(ctx, key, data, time.Hour); setErr != nil {
			fmt.Printf("s.users.GetDistrict: redis SET %q failed: %v\n", key, setErr)
		}
	} else {
		fmt.Printf("s.users.GetDistrict: json.Marshal(district) failed: %v\n", merr)
	}

	return district, nil
}

func (s *masterServiceImpl) GetSubDistrict(DistrictID int, ctx context.Context) ([]model.SubDistrict, error) {
	key := fmt.Sprintf("subdistrict:district:%d", DistrictID)

	subdistrict_redis_data, err := s.rc.Get(ctx, key)
	fmt.Println("redis data: ", subdistrict_redis_data)
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("redis GET %q: %w", key, err)
	}

	if len(subdistrict_redis_data) > 0 {
		var cached []model.SubDistrict
		if u_err := json.Unmarshal([]byte(subdistrict_redis_data), &cached); u_err == nil {
			fmt.Println("return cached data")
			return cached, nil
		}
	}

	subdistrict, dberr := s.repo.GetSubDistrict(DistrictID, ctx)
	if dberr != nil {
		return nil, fmt.Errorf("GetSubdistrict repo: %w", dberr)
	}
	if len(subdistrict) == 0 {
		fmt.Println("s.users.GetSubdistrict: No district ID")
		return nil, nil
	}

	if data, merr := json.Marshal(subdistrict); merr == nil {
		if setErr := s.rc.Set(ctx, key, data, time.Hour); setErr != nil {
			fmt.Printf("s.users.GetSubdistrict: redis SET %q failed: %v\n", key, setErr)
		}
	} else {
		fmt.Printf("s.users.GetSubdistrict: json.Marshal(district) failed: %v\n", merr)
	}

	return subdistrict, nil
}
