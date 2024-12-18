package main

import (
	"context"
	"fmt"
	"time"

	"github.com/caodonghui426/trpc-go-demo/proto/simplest"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	simplest.RegisterHelloWorldService(s, helloWorldImpl{})
	_ = s.Serve()
}

type helloWorldImpl struct{}

func (helloWorldImpl) Hello(ctx context.Context, req *simplest.HelloRequest) (*simplest.HelloResponse, error) {
	rsp := &simplest.HelloResponse{}
	rsp.Response = fmt.Sprintf("%s to you, too", req.Greeting)
	rsp.TimestampMsec = time.Now().UnixMilli()
	return rsp, nil
}
