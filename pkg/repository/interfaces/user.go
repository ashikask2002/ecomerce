package interfaces

import "github.com/ashikask2002/ecomerce.git/pkg/utils/models"

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockedStatus(email string) (bool, error)
	FindUserByEmail(email string) (models.UserSignInResponse, error)
}
