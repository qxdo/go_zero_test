package api

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"go_zero_test/inner/config"
	"go_zero_test/inner/handler"
	"go_zero_test/inner/svc"
	"net/http"
	"bytes"
	"os/exec"
)

var srv rest.Server

func init() {
	
	var c config.Config
	// 执行 ls 命令
	cmd := exec.Command("ls", "-la") // 使用 "-la" 选项以长格式列出所有文件，包括隐藏文件
	var out bytes.Buffer
	cmd.Stdout = &out // 将输出重定向到 out 变量

	// 运行命令
	err := cmd.Run()
	if err != nil {
		fmt.Println("执行 ls 命令失败:", err)
		return
	}

	// 输出 ls 命令的结果
	fmt.Print(out.String())
	
	conf.MustLoad("./etc/gozerotest-api.yaml", &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
