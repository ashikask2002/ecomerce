package usecase

import (
	"errors"
	"fmt"

	"github.com/ashikask2002/ecomerce.git/pkg/helper"
	"github.com/ashikask2002/ecomerce.git/pkg/repository/interfaces"
	usecase "github.com/ashikask2002/ecomerce.git/pkg/usecase/interfaces"
	"github.com/ashikask2002/ecomerce.git/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCaseImpl struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) usecase.UserUseCase {
	return &userUseCaseImpl{
		userRepo: repo,
	}
}

func (u *userUseCaseImpl) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	userExist := u.userRepo.CheckUserAvailability(user.Email)

	fmt.Println("user exist", userExist)
	// fmt.Println("user email", user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist , sign in")
	}
	fmt.Println(user)
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	// Hash password since detail are validated

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	//add userdetails to the database

	userData, err := u.userRepo.UserSignUp(user)
	//fmt.Println("userdata", userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	//jwt token String for user

	tokenString, err := helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	//copies all the details except the password of the user
	var UserDetails models.UserDetailsResponse
	err = copier.Copy(&UserDetails, &userData)
	//fmt.Prinln("userDetails:",userDetails)
	//fmt.Println("userData:", userData)

	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: UserDetails,
		Token: tokenString,
	}, nil
}

// log in
func (u *userUseCaseImpl) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	ok := u.userRepo.CheckUserAvailability(user.Email)

	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}
	isBlocked, err := u.userRepo.UserBlockedStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal error")
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}
	user_details, err := u.userRepo.FindUserByEmail(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = int(user_details.Id)
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal issue")

	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}
