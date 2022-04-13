package userRoutes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users/usersHandler"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users/usersRepository"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users/usersUsecase"
)

// UserRoutes - set routes for user endpoint
func UserRoutes(r *gin.Engine, dbSql *sql.DB) {

	userRepo := usersRepository.NewUserRepository(dbSql)
	userUC := usersUsecase.NewUserUsecase(userRepo)
	usersHandler.NewUserHTTPHandler(r, userUC)
}
