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
	Code       string    `gorm:"column:code;type:char(6);index:idx_code,unique;" json:"code"`
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

type Vendor struct {
	ID      uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name    string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	Address string    `gorm:"type:text;" json:"address"`
	Phone   string    `gorm:"type:varchar(25);" json:"phone"`
	Brand   string    `gorm:"type:varchar(255);" json:"brand"`
	common.WithTimestampsModel
}

type Item struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CityID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_city_id;not null" json:"city_id"`
	TypeID    uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_type_id;not null" json:"type_id"`
	VendorID  uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_vendor_id;not null" json:"vendor_id"`
	Name      string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`
	Address   string    `gorm:"type:text;not null;" json:"address"`
	Latitude  float64   `gorm:"type:decimal(10,8);" json:"latitude"`
	Longitude float64   `gorm:"type:decimal(10,8);" json:"longitude"`
	Position  uint8     `gorm:"type:smallint;not null;default:0" json:"position"`
	Width     float64   `gorm:"type:decimal(10,2);default:0.0" json:"width"`
	Height    float64   `gorm:"type:decimal(10,2);default:0.0" json:"height"`
	common.WithTimestampsModel
	City   City   `gorm:"foreignKey:CityID" json:"city"`
	Type   Type   `gorm:"foreignKey:TypeID" json:"type"`
	Vendor Vendor `gorm:"foreignKey:VendorID" json:"vendor"`
}
