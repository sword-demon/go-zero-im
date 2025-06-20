package svc

import (
	"github.com/sword-demon/go-zero-im/apps/user/api/internal/config"
	"github.com/sword-demon/go-zero-im/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
