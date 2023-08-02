package scheme

import (
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID       uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email    string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password *string        `gorm:"type:text" json:"password"`
	Roles    datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	common.WithTimestampsModel
}

type Province struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Code string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	common.WithTimestampsModel
}

type City struct {
	ID         uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProvinceID uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_province_id;not null" json:"province_id"`
	Code       string    `gorm:"column:code;type:char(4);index:idx_code,unique;" json:"code"`
	Name       string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	common.WithTimestampsModel
}

type Type struct {
	ID   uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	Icon string    `gorm:"column:icon;type:text;" json:"icon"`
	common.WithTimestampsModel
}
