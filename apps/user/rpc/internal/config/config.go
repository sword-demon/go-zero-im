package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// 定义配置属性

	Mysql struct {
		DataSource string
	}

	Cache cache.CacheConf

	Jwt struct {
		AccessSecret string
		AccessExpire int64
	}
}
