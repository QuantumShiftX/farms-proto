// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: manage.proto

package manageapiservice

import (
	"context"

	"github.com/QuantumShiftX/farms-proto/proto-gen-go/manage/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CollectionRewards         = v1.CollectionRewards
	DefaultVipInfoReply       = v1.DefaultVipInfoReply
	DepositRewards            = v1.DepositRewards
	GetAgreementReply         = v1.GetAgreementReply
	GetApkAddressReply        = v1.GetApkAddressReply
	GetDownloadAddrReply      = v1.GetDownloadAddrReply
	GetNotificationsListReply = v1.GetNotificationsListReply
	GetNotificationsListReq   = v1.GetNotificationsListReq
	ManageReply               = v1.ManageReply
	ManageReq                 = v1.ManageReq
	NotificationInfo          = v1.NotificationInfo
	RegistrationRewards       = v1.RegistrationRewards
	SendCaptchaReq            = v1.SendCaptchaReq
	SettingBaseInfoReply      = v1.SettingBaseInfoReply
	SettlementConfig          = v1.SettlementConfig
	VIPLevelInfo              = v1.VIPLevelInfo
	VipLevelDetail            = v1.VipLevelDetail
	VipLevelInfoMsgReply      = v1.VipLevelInfoMsgReply
	VipLevelInfoMsgReq        = v1.VipLevelInfoMsgReq

	ManageApiService interface {
		SendCaptcha(ctx context.Context, in *SendCaptchaReq, opts ...grpc.CallOption) (*ManageReply, error)
		GetAgreement(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*GetAgreementReply, error)
		GetApkAddress(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*GetApkAddressReply, error)
	}

	defaultManageApiService struct {
		cli zrpc.Client
	}
)

func NewManageApiService(cli zrpc.Client) ManageApiService {
	return &defaultManageApiService{
		cli: cli,
	}
}

func (m *defaultManageApiService) SendCaptcha(ctx context.Context, in *SendCaptchaReq, opts ...grpc.CallOption) (*ManageReply, error) {
	client := v1.NewManageApiServiceClient(m.cli.Conn())
	return client.SendCaptcha(ctx, in, opts...)
}

func (m *defaultManageApiService) GetAgreement(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*GetAgreementReply, error) {
	client := v1.NewManageApiServiceClient(m.cli.Conn())
	return client.GetAgreement(ctx, in, opts...)
}

func (m *defaultManageApiService) GetApkAddress(ctx context.Context, in *ManageReq, opts ...grpc.CallOption) (*GetApkAddressReply, error) {
	client := v1.NewManageApiServiceClient(m.cli.Conn())
	return client.GetApkAddress(ctx, in, opts...)
}
