package entity

import "time"

type UserEntity struct {
	// Khóa chính
	Id string `gorm:"column:id;primaryKey;type:char(32)"`
	// Email
	Email string `gorm:"column:email;index:idx_email;type:varchar(255);not null"`
	// Mật khẩu
	PasswordHashed string `gorm:"column:passwordHashed;type:varchar(255);not null"`
	// Thời gian tạo
	CreatedDate time.Time `gorm:"column:createdDate;type:datetime;not null;default:current_timestamp"`
	// Thời gian cập nhật
	UpdatedDate time.Time `gorm:"column:updatedDate;type:datetime"`
}

func (UserEntity) TableName() string {
	return "soraUsers"
}
