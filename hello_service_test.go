package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/trpcproto/trpchello"
	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/trpcproto/trpchello/trpc_hello_mock.go -package=trpchello -self_package=github.com/trpcproto/trpchello --source=stub/github.com/trpcproto/trpchello/trpc_hello.trpc.go

func Test_helloServiceImpl_HandleFirst(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helloServiceService := pb.NewMockHelloServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := helloServiceService.EXPECT().HandleFirst(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloReq) (*pb.HelloRsp, error) {
		s := &helloServiceImpl{}
		return s.HandleFirst(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.HelloReq
		rsp *pb.HelloRsp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloRsp
			var err error
			if rsp, err = helloServiceService.HandleFirst(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("helloServiceImpl.HandleFirst() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("helloServiceImpl.HandleFirst() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_helloServiceImpl_HandleSeconds(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	helloServiceService := pb.NewMockHelloServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := helloServiceService.EXPECT().HandleSeconds(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloSecondReq) (*pb.HelloSecondRsp, error) {
		s := &helloServiceImpl{}
		return s.HandleSeconds(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.HelloSecondReq
		rsp *pb.HelloSecondRsp
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloSecondRsp
			var err error
			if rsp, err = helloServiceService.HandleSeconds(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("helloServiceImpl.HandleSeconds() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("helloServiceImpl.HandleSeconds() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
