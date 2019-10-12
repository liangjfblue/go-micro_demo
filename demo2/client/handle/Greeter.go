package handle

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
)

type Greeter struct {
	client.Client
}

func (l *Greeter) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Printf("[Greeter] client request service: %s， method: %s\n", req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func LogWrap(c client.Client) client.Client {
	return &Greeter{c}
}
