package svc

import (
	"github.com/sword-demon/go-zero-im/apps/social/rpc/internal/config"
	"github.com/sword-demon/go-zero-im/apps/social/socialmodels"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	socialmodels.FriendsModel
	socialmodels.FriendRequestsModel
	socialmodels.GroupsModel
	socialmodels.GroupMembersModel
	socialmodels.GroupRequestsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		FriendsModel:        socialmodels.NewFriendsModel(conn, c.Cache),
		FriendRequestsModel: socialmodels.NewFriendRequestsModel(conn, c.Cache),
		GroupsModel:         socialmodels.NewGroupsModel(conn, c.Cache),
		GroupMembersModel:   socialmodels.NewGroupMembersModel(conn, c.Cache),
		GroupRequestsModel:  socialmodels.NewGroupRequestsModel(conn, c.Cache),
	}
}
