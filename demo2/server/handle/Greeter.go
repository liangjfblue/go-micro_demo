package handle

import (
	"context"
	proto "demo2/server/proto"
	"log"

	"github.com/micro/go-micro/server"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

/**
执行请求处理前的处理，一般用于熔断，限流，Filter，鉴权等
*/
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[logWrapper] server request: %v", req.Method())
		err := fn(ctx, req, rsp)
		return err
	}
}
