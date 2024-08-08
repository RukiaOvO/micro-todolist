package main

import (
	"fmt"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"micro-todolist/app/gateway/router"
	"micro-todolist/app/gateway/rpc"
	"micro-todolist/config"
	"time"
)

func main() {
	config.Init()
	rpc.InitRPC()

	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	webService := web.NewService(
		web.Name("httpService"),
		web.Address("localhost:4000"),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	_ = webService.Init()
	_ = webService.Run()
}
