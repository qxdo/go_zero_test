package api

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go_zero_test/inner/config"
	"go_zero_test/inner/handler"
	"go_zero_test/inner/svc"
	"net/http"
)

var srv rest.Server

func init() {
	var c config.Config
	conf.MustLoad("../etc/gozerotest-api.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
