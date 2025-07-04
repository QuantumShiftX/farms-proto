// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: dispatcher.proto

package dispatchertimer

import (
	"context"

	"github.com/QuantumShiftX/farms-proto/proto-gen-go/dispatcher/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CropReadyToHarvestReq           = v1.CropReadyToHarvestReq
	DispatcherReply                 = v1.DispatcherReply
	DispatcherReq                   = v1.DispatcherReq
	HeartbeatEventFastReq           = v1.HeartbeatEventFastReq
	HeartbeatEventSlowReq           = v1.HeartbeatEventSlowReq
	TriggerUserEventReq             = v1.TriggerUserEventReq
	UpdateCropStatusCheckReq        = v1.UpdateCropStatusCheckReq
	UpdateFortuneTreeStatusCheckReq = v1.UpdateFortuneTreeStatusCheckReq
	UpdateOnlineRewardTaskReq       = v1.UpdateOnlineRewardTaskReq
	UserFriendActionEventReq        = v1.UserFriendActionEventReq
	UserLoginEventReq               = v1.UserLoginEventReq
	UserOnlineDuration              = v1.UserOnlineDuration
	UserRechargeEventReq            = v1.UserRechargeEventReq
	UserRegistrationEventReq        = v1.UserRegistrationEventReq
	UserWithdrawEventReq            = v1.UserWithdrawEventReq

	DispatcherTimer interface {
		// 发财树状态检查
		CycleFertileTreeStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
		// 农场作物状态更新
		CycleCropStageUpdate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
		// 代理奖励结算
		CycleSettleAgentReward(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
		// 周期性更新基准汇率数据
		CycleUpdateBaseExchangeRate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
		// 周期性更新实时汇率数据
		CycleUpdateRealExchangeRate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
	}

	defaultDispatcherTimer struct {
		cli zrpc.Client
	}
)

func NewDispatcherTimer(cli zrpc.Client) DispatcherTimer {
	return &defaultDispatcherTimer{
		cli: cli,
	}
}

// 发财树状态检查
func (m *defaultDispatcherTimer) CycleFertileTreeStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	client := v1.NewDispatcherTimerClient(m.cli.Conn())
	return client.CycleFertileTreeStatusCheck(ctx, in, opts...)
}

// 农场作物状态更新
func (m *defaultDispatcherTimer) CycleCropStageUpdate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	client := v1.NewDispatcherTimerClient(m.cli.Conn())
	return client.CycleCropStageUpdate(ctx, in, opts...)
}

// 代理奖励结算
func (m *defaultDispatcherTimer) CycleSettleAgentReward(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	client := v1.NewDispatcherTimerClient(m.cli.Conn())
	return client.CycleSettleAgentReward(ctx, in, opts...)
}

// 周期性更新基准汇率数据
func (m *defaultDispatcherTimer) CycleUpdateBaseExchangeRate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	client := v1.NewDispatcherTimerClient(m.cli.Conn())
	return client.CycleUpdateBaseExchangeRate(ctx, in, opts...)
}

// 周期性更新实时汇率数据
func (m *defaultDispatcherTimer) CycleUpdateRealExchangeRate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	client := v1.NewDispatcherTimerClient(m.cli.Conn())
	return client.CycleUpdateRealExchangeRate(ctx, in, opts...)
}
