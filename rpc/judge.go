package rpc

import (
	"context"
	"github.com/l-f-h/common"
	grpclb "github.com/l-f-h/grpc-lb/etcdv3"
	"github.com/l-f-h/judge_server/conf"
	"github.com/l-f-h/judge_server/rpc/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

var conn *grpc.ClientConn

const (
	ServiceName string = "JudgeService"
)

func init() {
	var err error
	r := grpclb.NewResolver(conf.EtcdAddress, ServiceName)
	resolver.Register(r)
	url := r.Scheme() + "://authority/" + ServiceName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	conn, err = grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("Can not connect Service %s:%v", ServiceName, err)
	}
}

func Judge(ctx context.Context, req *pb_gen.JudgeRequest) (resp *pb_gen.JudgeResponse) {
	client := pb_gen.NewJudgeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
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
