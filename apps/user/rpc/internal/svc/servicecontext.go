package svc

import (
	"github.com/sword-demon/go-zero-im/apps/user/models"
	"github.com/sword-demon/go-zero-im/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// 添加用户模型，以便于在业务中使用
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 先创建好sql连接
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}
