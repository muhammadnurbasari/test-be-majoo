package users

import "github.com/muhammadnurbasari/test-be-majoo/models/loginModel"

type UsersUsecase interface {
	Login(myData loginModel.ReqLogin) (*loginModel.ResLogin, error)
}
