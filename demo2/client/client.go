package main

import (
	"context"
	"demo2/client/handle"
	"fmt"

	"github.com/micro/go-micro/registry"

	"github.com/micro/go-plugins/registry/etcdv3"

	proto "demo2/server/proto"

	"github.com/micro/go-micro"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://172.16.7.16:9002",
			"http://172.16.7.16:9004",
			"http://172.16.7.16:9006",
		}
	})

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(reg),
		micro.WrapClient(handle.LogWrap),
	)
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "liangjf"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
