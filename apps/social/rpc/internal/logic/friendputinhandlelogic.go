package logic

import (
	"context"

	"github.com/sword-demon/go-zero-im/apps/social/rpc/internal/svc"
	"github.com/sword-demon/go-zero-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// 获取好有申请记录
	// 验证是否有处理
	// 修改申请结果 > 通过[建立两条好友关系记录] -> 事务
	// 修改申请结果 > 拒绝[删除申请记录] -> 事务

	return &social.FriendPutInHandleResp{}, nil
}
