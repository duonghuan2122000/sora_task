package tenantrepository

import (
	"sorataskapi/internal/database"
	"sorataskapi/internal/entity"
	"time"
)

type TenantRepository interface {
	GetListByUser(userId string) ([]*entity.TenantEntity, error)
	Insert(tenantEntity entity.TenantEntity) error
	GetById(id string, userId string) (*entity.TenantEntity, error)
}

type tenantRepo struct{}

func NewTenantRepository() TenantRepository {
	return &tenantRepo{}
}

func (repo *tenantRepo) GetListByUser(userId string) ([]*entity.TenantEntity, error) {
	var tenants []*entity.TenantEntity
	result := database.MysqlConnect.Raw("SELECT id, code, name, status, createdDate, updatedDate, createdBy, updatedBy FROM soraTenants t WHERE t.id IN (SELECT tenantId FROM soraTenantUsers WHERE userId = ?)", userId).Scan(&tenants)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return []*entity.TenantEntity{}, nil
	}
	return tenants, nil
}

func (repo *tenantRepo) Insert(tenantEntity entity.TenantEntity) error {
	tenantEntity.CreatedDate = time.Now().UTC()
	return database.MysqlConnect.Create(&tenantEntity).Error
}

func (repo *tenantRepo) GetById(id string, userId string) (*entity.TenantEntity, error) {
	var tenant entity.TenantEntity
	result := database.MysqlConnect.Raw("SELECT id, code, name, status, createdDate, updatedDate, createdBy, updatedBy FROM soraTenants t WHERE t.id = ? AND t.id IN (SELECT tenantId FROM soraTenantUsers WHERE userId = ?) LIMIT 1", id, userId).Scan(&tenant)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &tenant, nil
}
