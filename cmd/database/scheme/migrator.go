package scheme

import "gorm.io/gorm"

func prepareTable() []interface{} {
	return []interface{}{
		&User{},
		&Province{},
		&City{},
		&Type{},
		&Vendor{},
	}
}
func Migrate(db *gorm.DB) {
	tables := prepareTable()
	db.AutoMigrate(tables...)
}
