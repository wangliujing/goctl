package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
    "github.com/wangliujing/foundation-framework/system"
	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
    logx.MustSetup(c.Log)
	server := rest.MustNewServer(c.RestConf)

	ctx := svc.NewServiceContext(&c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	system.Start(server)
}
