package scheme

import "gorm.io/gorm"

func prepareTable() []interface{} {
	return []interface{}{
		&User{},
		&Province{},
		&City{},
		&MediaType{},
		&Vendor{},
		&Item{},
		&Project{},
		&ProjectItem{},
		&ItemImage{},
	}
}
func Migrate(db *gorm.DB) {
	tables := prepareTable()
	db.AutoMigrate(tables...)
}

func Drop(db *gorm.DB) {
	tables := prepareTable()
	db.Migrator().DropTable(tables...)
}
