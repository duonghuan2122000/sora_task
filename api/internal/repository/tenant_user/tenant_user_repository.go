package tenantuserrepository

import (
	"sorataskapi/internal/database"
	"sorataskapi/internal/entity"
)

type TenantUserRepository interface {
	Insert(tenantUserEntity entity.TenantUserEntity) error
}

type tenantUserRepo struct {
}

func NewTenantUserRepository() TenantUserRepository {
	return &tenantUserRepo{}
}

func (repo *tenantUserRepo) Insert(tenantUserEntity entity.TenantUserEntity) error {
	return database.MysqlConnect.Create(&tenantUserEntity).Error
}
