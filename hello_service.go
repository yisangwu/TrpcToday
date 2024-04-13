package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/trpcproto/trpchello"
)

type helloServiceImpl struct {
	pb.UnimplementedHelloService
}

func (s *helloServiceImpl) HandleFirst(
	ctx context.Context,
	req *pb.HelloReq,
) (*pb.HelloRsp, error) {

	message := req.Message
	if message == "" {
		message = "empty message"
	}
	rsp := &pb.HelloRsp{}
	rsp.Code = 100
	rsp.Message = fmt.Sprintf("return client message:%s", message)
	rsp.Timestamp = time.Now().Unix()
	return rsp, nil
}

func (s *helloServiceImpl) HandleSeconds(ctx context.Context, req *pb.HelloSecondReq) (*pb.HelloSecondRsp, error) {
	rsp := &pb.HelloSecondRsp{}
	rsp.Code = 200
	rsp.Message = "second success"
	rsp.Timestamp = time.Now().Unix()
	return rsp, nil
}
