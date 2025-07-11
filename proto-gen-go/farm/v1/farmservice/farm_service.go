// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: farm.proto

package farmservice

import (
	"context"

	"github.com/QuantumShiftX/farms-proto/proto-gen-go/farm/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CropsInfoListMsgReply     = v1.CropsInfoListMsgReply
	CropsInfoListMsgReq       = v1.CropsInfoListMsgReq
	CropsProductInfo          = v1.CropsProductInfo
	CropsProductInfoMsgReply  = v1.CropsProductInfoMsgReply
	CropsProductInfoMsgReq    = v1.CropsProductInfoMsgReq
	FarmReply                 = v1.FarmReply
	FarmReq                   = v1.FarmReq
	FarmsStoreInfoMsgReply    = v1.FarmsStoreInfoMsgReply
	FarmsStoreInfoMsgReq      = v1.FarmsStoreInfoMsgReq
	GenerateFarmsNameMsgReply = v1.GenerateFarmsNameMsgReply
	GenerateFarmsNameMsgReq   = v1.GenerateFarmsNameMsgReq
	MsgReply                  = v1.MsgReply
	MsgReq                    = v1.MsgReq
	StoreInfo                 = v1.StoreInfo
	StoreProductInfo          = v1.StoreProductInfo
	StoreProductInfoMsgReply  = v1.StoreProductInfoMsgReply
	StoreProductInfoMsgReq    = v1.StoreProductInfoMsgReq
	StoreProductListMsgReply  = v1.StoreProductListMsgReply
	StoreProductListMsgReq    = v1.StoreProductListMsgReq

	FarmService interface {
		Test(ctx context.Context, in *FarmReq, opts ...grpc.CallOption) (*FarmReply, error)
	}

	defaultFarmService struct {
		cli zrpc.Client
	}
)

func NewFarmService(cli zrpc.Client) FarmService {
	return &defaultFarmService{
		cli: cli,
	}
}

func (m *defaultFarmService) Test(ctx context.Context, in *FarmReq, opts ...grpc.CallOption) (*FarmReply, error) {
	client := v1.NewFarmServiceClient(m.cli.Conn())
	return client.Test(ctx, in, opts...)
}
