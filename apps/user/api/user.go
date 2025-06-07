package main

import (
	"flag"
	"fmt"
	"github.com/sword-demon/go-zero-im/pkg/ctxdata/resultx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/sword-demon/go-zero-im/apps/user/api/internal/config"
	"github.com/sword-demon/go-zero-im/apps/user/api/internal/handler"
	"github.com/sword-demon/go-zero-im/apps/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dev/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 使用 自定义的http 响应
	httpx.SetErrorHandlerCtx(resultx.ErrHandler(c.Name))
	httpx.SetOkHandler(resultx.OKHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
