package entity

type TenantUserEntity struct {
	TenantId string `gorm:"column:tenantId;primaryKey;type:char(32)"`
	UserId   string `gorm:"column:userId;primaryKey;type:char(32)"`
}

func (TenantUserEntity) TableName() string {
	return "soraTenantUsers"
}
