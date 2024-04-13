package main

import (
	pb "github.com/trpcproto/trpchello"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloServiceService(s.Service("trpchello.HelloService"), &helloServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
