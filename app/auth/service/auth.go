package service

import (
	"errors"

	"github.com/bagusyanuar/go-yousee/app/auth/repositories"
	"github.com/bagusyanuar/go-yousee/app/auth/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/config"
	"github.com/bagusyanuar/go-yousee/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	AuthService interface {
		SignIn(request request.AuthRequest) (string, error)
	}

	Auth struct {
		jwtConfig      config.JWT
		authRepository repositories.AuthRepository
	}
)

// SignIn implements AuthService.
func (svc *Auth) SignIn(request request.AuthRequest) (string, error) {
	entity := model.User{
		Username: request.Username,
		Password: &request.Password,
	}

	user, err := svc.authRepository.SignIn(entity)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return "", common.ErrUserNotFound
		}
		return "", err
	}

	errComparePassword := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	if errComparePassword != nil {
		return "", common.ErrPasswordNotMatch
	}

	jwtSign := common.JWTSignReturn{
		UserID: user.ID,
	}
	return common.CreateAccessToken(&svc.jwtConfig, &jwtSign)
}

func NewAuth(authRepo repositories.AuthRepository, jwtCfg config.JWT) AuthService {
	return &Auth{
		authRepository: authRepo,
		jwtConfig:      jwtCfg,
	}
}
