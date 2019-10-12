package main

import (
	"demo2/server/config"
	"demo2/server/handle"
	"log"

	proto "demo2/server/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	_ "github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = config.InitConfig()
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(reg),
	)

	// 初始化服务
	// Init will parse the command line flags.
	service.Init(
		micro.WrapHandler(handle.LogWrapper),
	)

	// Register handler
	if err := proto.RegisterGreeterHandler(service.Server(), new(handle.Greeter)); err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("go micro start...")
	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
		return
	}
}
