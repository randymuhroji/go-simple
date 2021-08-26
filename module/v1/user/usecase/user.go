package usecase

import (
	"errors"
	"go-simple/config"
	"go-simple/model"
	"go-simple/module/v1/user/repo"
	"go-simple/utl/password"
	"strings"
)

func UserList(conf config.Configuration) (users []model.UserView, err error) {
	db := conf.MysqlDB
	return repo.GetUserList(db)
}

// get user detail
func UserDetail(conf config.Configuration, userId int) (users model.UserView, err error) {
	db := conf.MysqlDB
	return repo.GetUserDetail(db, userId)
}

// register new user
func UserRegister(conf config.Configuration, usr *model.User) (user model.UserView, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Password) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	hash, err := password.Encrypt(usr.Password)
	if err != nil {
		return user, err
	}

	payload.Password = string(hash)
	payload.Name = usr.Name
	payload.Email = usr.Email

	sqlResult, err := repo.RegisterUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	id, _ := sqlResult.LastInsertId()

	return model.UserView{
		Id:    int(id),
		Name:  *&usr.Name,
		Email: *&usr.Email,
	}, nil
}

// update user
func UserUpdate(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	if strings.TrimSpace(usr.Name) == "" || strings.TrimSpace(usr.Email) == "" {
		return user, errors.New("data can not null")
	}

	payload.Name = usr.Name
	payload.Email = usr.Email
	payload.Id = usr.Id

	_, err = repo.UpdateUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}

// delete user
func UserDelete(conf config.Configuration, usr *model.User) (user model.User, err error) {
	var (
		payload = model.UserPayload{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	payload.Id = usr.Id
	payload.Deleted = 1

	_, err = repo.DeleteUser(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *usr, nil
}
