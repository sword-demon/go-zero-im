package logic

import (
	"context"
	"database/sql"
	googleErr "errors"
	"github.com/pkg/errors"
	"github.com/sword-demon/go-zero-im/apps/social/socialmodels"
	"github.com/sword-demon/go-zero-im/pkg/constants"
	"github.com/sword-demon/go-zero-im/pkg/xerr"
	"time"

	"github.com/sword-demon/go-zero-im/apps/social/rpc/internal/svc"
	"github.com/sword-demon/go-zero-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// 申请人是否与目标是好有关系 申请人 id 和目标 id 进行查询    uid, fid
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && !googleErr.Is(err, socialmodels.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friends by uid and fid err %v req %v", err, in)
	}
	if friends != nil {
		return &social.FriendPutInResp{}, err
	}
	// 是否已经有过申请,并且申请是不成功的,没有完成的
	friendReqs, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && !googleErr.Is(err, socialmodels.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend reqs by rid and uid err %v, req %v", err, in)
	}
	if friendReqs != nil {
		return &social.FriendPutInResp{}, err
	}
	// 创建申请记录
	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &socialmodels.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})

	// 异常处理
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert friendRequest err %v, req %v", err, in)
	}

	return &social.FriendPutInResp{}, nil
}
