package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		SignIn(entity model.User) (*model.User, error)
	}

	Auth struct {
		database *gorm.DB
	}
)

// SignIn implements AuthRepository.
func (r *Auth) SignIn(entity model.User) (*model.User, error) {
	user := new(model.User)
	if err := r.database.Where("username = ?", entity.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuth(db *gorm.DB) AuthRepository {
	return &Auth{
		database: db,
	}
}
