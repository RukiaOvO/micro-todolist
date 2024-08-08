package main

import (
	"fmt"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro-todolist/app/user/repository/db/dao"
	"micro-todolist/app/user/repository/service"
	"micro-todolist/config"
	"micro-todolist/idl/pb"
)

func main() {
	config.Init()
	dao.InitDB()

	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)
	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	_ = microService.Run()
}
