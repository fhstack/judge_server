package rpc

import (
	"context"
	"github.com/l-f-h/common"
	"github.com/l-f-h/judge_server/conf"
	"github.com/l-f-h/judge_server/rpc/generated"

	"google.golang.org/grpc"
	"log"
)

var conn *grpc.ClientConn

const (
	ServiceName string = "Judge"
)

func init() {
	var err error
	address := conf.RpcIp + ":" + conf.RpcPort
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can not connect Service %s:%v", ServiceName, err)
	}
}

func Judge(ctx context.Context, req *pb_gen.JudgeRequest) (resp *pb_gen.JudgeResponse) {
	client := pb_gen.NewJudgeServiceClient(conn)
	respJudge, err := client.Judge(ctx, req)
	if err != nil {
		log.Printf("Rpc Judge failed: %v", err)
		resp = &pb_gen.JudgeResponse{}
		resp.ResCode = common.RpcFailed
		return resp
	} else {
		resp = respJudge
	}
	return
}
