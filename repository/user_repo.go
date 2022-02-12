package repository

import (
	"context"
	"lienquanMess/model"
	"lienquanMess/model/req"
)

type UserRepo interface {
	CheckLogin(context context.Context, loginReq req.ReqSignIp) (model.User, error)
	SaveUser(context context.Context, user model.User, userInfo model.AllUser, userProfile model.UserProfile) (model.User, error)
	SelectUserProfileById(context context.Context, userId string) (model.UserProfile, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
	UpdateLineUser(context context.Context, user model.UserProfile) (model.UserProfile, error)
	SelectAllUserById(context context.Context, start string, userId string) ([]model.AllUser, error)
}
