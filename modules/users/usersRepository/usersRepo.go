package usersRepository

import (
	"database/sql"
	"errors"

	"github.com/muhammadnurbasari/test-be-majoo/models/loginModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/users"
)

type sqlRepository struct {
	Conn *sql.DB
}

//NewUserRepository - will create object that represent that users.UserRepository interface
func NewUserRepository(Conn *sql.DB) users.UsersRepository {
	return &sqlRepository{Conn}
}

//RetrieveUserByUsername - get data for users ( with parameter username )
func (config *sqlRepository) RetrieveUserByUsername(username string) (*loginModel.Users, error) {
	query := "SELECT id,name,user_name,password FROM users WHERE user_name = ?"

	stat, errPrepare := config.Conn.Prepare(query)
	if errPrepare != nil {
		return nil, errors.New("RetrieveUserByUsername errPrepare = " + errPrepare.Error())
	}
	defer stat.Close()

	var resp loginModel.Users
	errQueryRow := stat.QueryRow(username).Scan(&resp.ID, &resp.Name, &resp.Username, &resp.Password)
	if errQueryRow != nil {
		if errQueryRow == sql.ErrNoRows {
			// return error string "1" for is not exist username
			return nil, errors.New("1")
		}

		return nil, errors.New("RetrieveUserByUsername errQueryRow = " + errQueryRow.Error())
	}

	return &resp, nil
}
