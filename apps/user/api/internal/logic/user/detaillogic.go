package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/sword-demon/go-zero-im/apps/user/rpc/user"
	"github.com/sword-demon/go-zero-im/pkg/ctxdata"

	"github.com/sword-demon/go-zero-im/apps/user/api/internal/svc"
	"github.com/sword-demon/go-zero-im/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDetailLogic 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	// go-zero 处理完信息后会把 token 解析的数据传入 context 中

	uid := ctxdata.GetUID(l.ctx)
	userInfoResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{Id: uid})
	if err != nil {
		return nil, err
	}

	var res types.UserInfoResp
	copier.Copy(&res, userInfoResp)

	return &res, nil
}
