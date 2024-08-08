package service

import (
	"context"
	"errors"
	"micro-todolist/app/user/repository/db/dao"
	"micro-todolist/app/user/repository/db/model"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func GetUserSrvHungury() *UserSrv {
	if UserSrvIns == nil {
		return new(UserSrv)
	}
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	resp.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		resp.Code = e.Error
		return errors.New("user not found")
	}
	if !user.CheckPassWord(req.Password) {
		resp.Code = e.Error
		return errors.New("incorrect password")
	}

	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	resp.Code = e.Success

	if req.Password != req.PasswordConfirm {
		resp.Code = e.Error
		return errors.New("incorrect password confirm")
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID > 0 {
		resp.Code = e.Error
		return errors.New("user already exists")
	}
	user = &model.User{
		UserName: req.UserName,
	}
	if err = user.SetPassWord(req.Password); err != nil {
		resp.Code = e.Error
		return
	}
	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}
