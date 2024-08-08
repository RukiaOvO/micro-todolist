package rpc

import (
	"go-micro.dev/v4"
	"micro-todolist/idl/pb"
)

var (
	UserService pb.UserService
)

func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService
}
