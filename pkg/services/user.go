package services

import (
	"moment/pkg/db"

	"gitlab.com/knopkalab/go/logger"
)

type User interface {
	CreateUser(resp *UserUpdateRequest) error
	GetUserByID(id int64, ctx db.Context) (*db.User, error)
	GetUserByPhone(phone string, ctx db.Context) (*db.User, error)
	UpdateUserByID(id int64, ctx db.Context, newData *UserUpdateRequest) (*db.User, error)
	UpdateUser(ctx db.Context, newData *UserUpdateRequest, user *db.User) (*db.User, error)
}

type UserService struct {
	db  db.DB
	log logger.Logger
}

type UserUpdateRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func (r *UserUpdateRequest) fillTo(u *db.User) {
	u.Phone = r.Phone
	u.Code = r.Code
}
func (r *UserUpdateRequest) toFullUser() *db.User {
	u := &db.User{
		Phone: r.Phone,
		Code:  r.Code,
	}
	return u
}

func (a *UserService) CreateUser(resp *UserUpdateRequest) error {
	user := resp.toFullUser()
	return a.db.UserSave(user)
}

func (a *UserService) GetUserByID(id int64, ctx db.Context) (*db.User, error) {
	return a.db.GetUserByID(id, ctx)
}

func (a *UserService) GetUserByPhone(phone string, ctx db.Context) (*db.User, error) {
	return a.db.GetUserByPhone(phone, ctx)
}

func (a *UserService) UpdateUserByID(id int64, ctx db.Context, newData *UserUpdateRequest) (*db.User, error) {
	user, err := a.db.GetUserByID(id, ctx)
	if err != nil {
		return nil, err
	}
	newData.fillTo(user)
	if err = a.db.UserSave(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (a *UserService) UpdateUser(ctx db.Context, newData *UserUpdateRequest, user *db.User) (*db.User, error) {
	newData.fillTo(user)
	if err := a.db.UserSave(user); err != nil {
		return nil, err
	}
	return user, nil
}

func NewUser(log logger.Logger, db db.DB) User {
	return &UserService{db: db, log: log}
}
