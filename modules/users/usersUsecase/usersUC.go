package usersUsecase

import (
	"errors"

	"github.com/muhammadnurbasari/test-be-majoo/helpers/validation"
	"github.com/muhammadnurbasari/test-be-majoo/middlewares/jwt"
	"github.com/muhammadnurbasari/test-be-majoo/models/loginModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users"
)

type usersUsecase struct {
	usersRepository users.UsersRepository
}

//NewUserUsecase - will create new an userUsecase object representation of users.UserUsecase interface
func NewUserUsecase(userRepo users.UsersRepository) users.UsersUsecase {
	return &usersUsecase{
		usersRepository: userRepo,
	}
}

//Login - use case for handle user login
func (userUC *usersUsecase) Login(myData loginModel.ReqLogin) (*loginModel.ResLogin, error) {

	// get user by username
	result, err := userUC.usersRepository.RetrieveUserByUsername(myData.Username)

	if err != nil {
		if err.Error() == "1" {
			return nil, errors.New("Wrong Username")
		}
		return nil, err
	}

	// check password
	checkPassword := validation.CheckPasswordHash(myData.Password, result.Password)
	if checkPassword == false {
		//wrong password
		return nil, errors.New("Wrong Password")
	}

	//generate token jwt
	myToken, err := jwt.GenerateTokenJwt(result.ID)
	if err != nil {
		return nil, err
	}

	if myToken == "" {
		return nil, errors.New("Failed Generate Token")
	}

	resp := loginModel.ResLogin{
		ID:       result.ID,
		Name:     result.Name,
		Username: result.Username,
		Token:    myToken,
	}

	return &resp, nil

}
