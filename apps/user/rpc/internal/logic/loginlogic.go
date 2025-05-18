package logic

import (
	"context"
	"errors"
	"github.com/sword-demon/go-zero-im/apps/user/models"
	"github.com/sword-demon/go-zero-im/pkg/ctxdata"
	"github.com/sword-demon/go-zero-im/pkg/encrypt"
	"time"

	"github.com/sword-demon/go-zero-im/apps/user/rpc/internal/svc"
	"github.com/sword-demon/go-zero-im/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneNotRegister  = errors.New("手机号码没有注册过")
	ErrUserPasswordError = errors.New("用户名或密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 1. 根据手机号码查询用户是否存在
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return nil, ErrPhoneNotRegister
		}
		return nil, err
	}

	// 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, ErrUserPasswordError
	}

	now := time.Now().Unix()
	expire := l.svcCtx.Config.Jwt.AccessExpire + now
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, expire, userEntity.Id)
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Token:  token,
		Id:     userEntity.Id,
		Expire: expire,
	}, nil
}
