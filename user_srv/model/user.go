package model

import (
	"gorm.io/gorm"
	"time"
)

/* baseModel is a base model that contains common fields for all models.
* It includes an ID, CreatedAt, UpdatedAt, and DeletedAt fields.
 */
type baseModel struct {
	ID        int32     `grom:"primary_key"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt
	isDeleted bool
}

type User struct {
	baseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;tyoe:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Nickname string     `gorm:"type:varchar(20);not null"`
	Birthday *time.Time `gorm:"type:datetime;"`
	Genders  string     `gorm:"column:gender;default:male;type:varchar(6); comment:'female表示女 male表示男'"`
	Role     int        `gorm:"column:role;default:1; comment:'1表示普通用户 2表示管理员'"`
}
