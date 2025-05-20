package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sword-demon/go-zero-im/apps/user/models"
	"github.com/sword-demon/go-zero-im/pkg/ctxdata"
	"github.com/sword-demon/go-zero-im/pkg/encrypt"
	"github.com/sword-demon/go-zero-im/pkg/wuid"
	"github.com/sword-demon/go-zero-im/pkg/xerr"
	"time"

	"github.com/sword-demon/go-zero-im/apps/user/rpc/internal/svc"
	"github.com/sword-demon/go-zero-im/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegister = xerr.New(xerr.RequestParamError, "手机号码已注册过")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 1. 验证用户是否注册过，根据手机号码验证
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	logx.Error("err: ", err)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "register to find user by phone: err: %v, req: %v", err, in)
	}

	if userEntity != nil {
		return nil, errors.WithStack(ErrPhoneIsRegister)
	}

	// 定义用户数据
	userEntity = &models.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex), // 值类型
			Valid: true,          // 是否写入
		},
	}

	// 处理密码
	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "gen hash password: err: %v", err)
		}

		userEntity.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}

	_, err = l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "insert user to db: err: %v, data: %v", err, userEntity)
	}

	// 生成token
	now := time.Now().Unix() // 时间戳 秒为单位
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "get jwt token err: %v", err)
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
