package usersHandler

import (
	"net/http"

	"github.com/muhammadnurbasari/test-be-majoo/middlewares/auth"
	"github.com/muhammadnurbasari/test-be-majoo/models/loginModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users"

	"github.com/gin-gonic/gin"
)

//userHandler - representation users.UserUsecase
type userHandler struct {
	userUsecase users.UsersUsecase
}

//NewUserHTTPHandler - will create new an userUsecase object representation of userHandler interface
func NewUserHTTPHandler(r *gin.Engine, userUsecase users.UsersUsecase) {
	handler := userHandler{
		userUsecase: userUsecase,
	}

	authorized := r.Group("/user")

	{
		authorized.POST("/login", auth.BasicAuth, handler.LoginUser)
	}

}

// Login godoc
// @Summary Login by username and password, please try using POSTMAN
// @Description Login endpoint to get Authorization token
// @ID login-username-password
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param authinfo body loginModel.ReqLogin true "Auth info"
// @Success 200 {object} userDTO.ResLoginResult
// @Failure 404 {object} userDTO.ResNotFound
// @Router /user/login [post]
func (handler *userHandler) LoginUser(c *gin.Context) {
	var req loginModel.ReqLogin
	errBind := c.BindJSON(&req)

	if errBind != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusBadRequest,
				"messages": "Please Contact Admin !!!, " + errBind.Error(),
			},
		)
		return
	}

	res, err := handler.userUsecase.Login(req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":   http.StatusNotFound,
				"messages": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":   http.StatusOK,
			"messages": "Login Success",
			"result":   &res,
		},
	)
}
