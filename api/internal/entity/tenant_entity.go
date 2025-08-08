package entity

import "time"

type TenantEntity struct {
	// Khóa chính
	Id string `gorm:"column:id;primaryKey;type:char(32)"`
	// Mã tenant
	Code string `gorm:"column:code;index:idx_code;type:varchar(255);not null"`
	// Tên tenant
	Name string `gorm:"column:name;type:varchar(255);not null"`
	// Trạng thái tenant
	Status TenantStatus `gorm:"column:status;type:varchar(255);not null"`
	// Thời gian tạo
	CreatedDate time.Time `gorm:"column:createdDate;type:datetime;not null;default:current_timestamp"`
	// Thời gian cập nhật
	UpdatedDate *time.Time `gorm:"column:updatedDate;type:datetime"`
}

func (TenantEntity) TableName() string {
	return "soraTenants"
}

type TenantStatus string

const (
	// Tenant đang hoạt động
	TenantActive TenantStatus = "ACTIVE"
	// Tenant ngừng hoạt động
	TenantDisable TenantStatus = "DISABLE"
)
