package userDTO

import (
	"github.com/muhammadnurbasari/test-be-majoo/models/loginModel"
)

// ResLoginResult - response example login module
type ResLoginResult struct {
	Status    int                 `json:"status" example:"200"`
	Messagges string              `json:"messages" example:"Login Success"`
	Result    loginModel.ResLogin `json:"result"`
}

//
type ResNotFound struct {
	Status    int    `json:"status" example:"404"`
	Messagges string `json:"messages" example:"Wrong Username"`
}
