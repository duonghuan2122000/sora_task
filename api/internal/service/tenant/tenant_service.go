package tenantservice

import (
	"context"
	"crypto/rand"
	"math/big"
	"sorataskapi/internal/entity"
	basemodel "sorataskapi/internal/model/base"
	tenantmodel "sorataskapi/internal/model/tenant"
	usermodel "sorataskapi/internal/model/user"
	tenantrepository "sorataskapi/internal/repository/tenant"
	tenantuserrepository "sorataskapi/internal/repository/tenant_user"
	"strings"

	"github.com/google/uuid"
)

type TenantService interface {
	GetListByUser(userId string) ([]*entity.TenantEntity, error)
	Create(ctx context.Context, payload tenantmodel.CreateTenantReqDto) (*tenantmodel.CreateTenantResDto, error)
	GetCurrent(ctx context.Context) (*entity.TenantEntity, error)
	Select(ctx context.Context, payload tenantmodel.SelectTenantReqDto) error
}

type tenantService struct {
	tenantRepo     tenantrepository.TenantRepository
	tenantUserRepo tenantuserrepository.TenantUserRepository
}

func NewTenantService(tenantRepo tenantrepository.TenantRepository, tenantUserRep tenantuserrepository.TenantUserRepository) TenantService {
	return &tenantService{tenantRepo: tenantRepo, tenantUserRepo: tenantUserRep}
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func (tenantSvc *tenantService) GetListByUser(userId string) ([]*entity.TenantEntity, error) {
	return tenantSvc.tenantRepo.GetListByUser(userId)
}

func (tenantSvc *tenantService) Create(ctx context.Context, payload tenantmodel.CreateTenantReqDto) (*tenantmodel.CreateTenantResDto, error) {
	currentUserDto, ok := usermodel.FromContext(ctx)
	if !ok {
		return nil, &basemodel.LogicError{
			Code:    "401",
			Message: "Unauthorized",
		}
	}

	tenantCode, err := randomString(8)
	if err != nil {
		return nil, err
	}
	tenantEntity := entity.TenantEntity{
		Id:        strings.ReplaceAll(uuid.New().String(), "-", ""),
		Code:      tenantCode,
		Name:      payload.Name,
		Status:    entity.TenantActive,
		CreatedBy: currentUserDto.UserId,
	}
	err = tenantSvc.tenantRepo.Insert(tenantEntity)
	if err != nil {
		return nil, err
	}
	tenantUserEntity := entity.TenantUserEntity{
		TenantId: tenantEntity.Id,
		UserId:   currentUserDto.UserId,
	}
	err = tenantSvc.tenantUserRepo.Insert(tenantUserEntity)
	if err != nil {
		return nil, err
	}

	return &tenantmodel.CreateTenantResDto{
		Id:   tenantEntity.Id,
		Code: tenantEntity.Code,
		Name: tenantEntity.Name,
	}, nil
}

func (tenantSvc *tenantService) GetCurrent(ctx context.Context) (*entity.TenantEntity, error) {
	currentUserDto, ok := usermodel.FromContext(ctx)

	if !ok {
		return nil, &basemodel.LogicError{
			Code:    "401",
			Message: "Unauthorized",
		}
	}

	tenantId, okTenant := tenantmodel.FromContext(ctx)
	if !okTenant {
		return nil, &basemodel.LogicError{
			Code:    "204",
			Message: "Chưa chọn tenant",
		}
	}
	result, err := tenantSvc.tenantRepo.GetById(tenantId, currentUserDto.UserId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, &basemodel.LogicError{
			Code:    "204",
			Message: "Chưa chọn tenant",
		}
	}
	return result, nil
}

func (tenantSvc *tenantService) Select(ctx context.Context, payload tenantmodel.SelectTenantReqDto) error {
	return nil
}
