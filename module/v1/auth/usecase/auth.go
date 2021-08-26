package usecase

import (
	"errors"
	"go-simple/config"
	"go-simple/model"
	userRepository "go-simple/module/v1/user/repo"
	"go-simple/utl/jwt"
	"go-simple/utl/password"
)

func Login(cnf config.Configuration, p *model.LoginPayload) (token model.AuthAccess, err error) {

	// validation of null data
	if p.Email == "" || p.Password == "" {
		err = errors.New("email password can not null")
		return
	}

	// find user
	getUser, err := userRepository.GetUserDetailByParam(cnf.MysqlDB, "user_email", p.Email)
	if err != nil {
		return
	}

	if err = password.Decrypt([]byte(getUser.Password), p.Password); err != nil {
		return
	}

	return jwt.Generate(getUser.Email)
}
