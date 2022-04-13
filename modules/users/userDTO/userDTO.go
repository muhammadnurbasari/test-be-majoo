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

// ForbiddenRes - response forbidden request
type ForbiddenRes struct {
	Status    int    `json:"status" example:"403"`
	Messagges string `json:"messages" example:"incorrect Email or Password"`
}

// ResRegister - response register account
type ResRegister struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"Verification code send to your email, Please Verification On 2 Minute !"`
	UserCode  string `json:"user_code" example:""`
}

//  ResVerificationOtp
type ResVerificationOtp struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"Verification Success, Please Login !"`
}

//  ResResendOtp
type ResResendOtp struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"Resend Success, Please Verification On 2 Minute !"`
}

//
type ResNotFound struct {
	Status    int    `json:"status" example:"404"`
	Messagges string `json:"messages" example:"Please Registration Your Account !"`
}

type ResOK struct {
	Status    int    `json:"status" example:"200"`
	Messagges string `json:"messages" example:"success"`
}
