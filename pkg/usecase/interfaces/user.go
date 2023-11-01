package interfaces

import "github.com/ashikask2002/ecomerce.git/pkg/utils/models"

type UserUseCase interface {
	UserSignUp(User models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
}
