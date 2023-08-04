package scheme

import "gorm.io/gorm"

func prepareTable() []interface{} {
	return []interface{}{
		&User{},
		&Province{},
		&City{},
		&Type{},
		&Vendor{},
		&Item{},
	}
}
func Migrate(db *gorm.DB) {
	tables := prepareTable()
	db.AutoMigrate(tables...)
}
