package users

import "github.com/muhammadnurbasari/test-be-majoo/models/loginModel"

type UsersRepository interface {
	RetrieveUserByUsername(username string) (*loginModel.Users, error)
}
