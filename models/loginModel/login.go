package loginModel

// manage login model
type (

	//ReqLogin - object for mapping request json login
	ReqLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//ResLogin - object response json login
	ResLogin struct {
		ID       int64  `json:"userid"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}
)

// manage user model
type (
	// Users - object for mapping response data users
	Users struct {
		ID       int64  `json:"userid"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
