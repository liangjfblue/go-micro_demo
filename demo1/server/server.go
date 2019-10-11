package main

import (
	"context"
	"log"

	proto "go-micro_demo/server/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	_ "github.com/micro/go-plugins/registry/etcdv3"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://172.16.7.16:9002",
			"http://172.16.7.16:9004",
			"http://172.16.7.16:9006",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(reg),
	)

	// 初始化服务
	// Init will parse the command line flags.
	service.Init()

	// Register handler
	if err := proto.RegisterGreeterHandler(service.Server(), new(Greeter)); err != nil {
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
