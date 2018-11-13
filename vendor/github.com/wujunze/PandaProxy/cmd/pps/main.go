package main

import (
	"fmt"
	"github.com/phayes/freeport"
	"github.com/wujunze/PandaProxy/PandaProxy"
	"github.com/wujunze/PandaProxy/cmd"
	"log"
	"net"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 服务端监听端口随机生成
	port, err := freeport.GetFreePort()
	if err != nil {
		// 随机端口失败就采用 7448
		port = 7448
	}
	// 默认配置
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		// 密码随机生成
		Password: PandaProxy.RandPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 server 端并监听
	lsServer, err := PandaProxy.NewLsServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
密码 password：
%s
	`, listenAddr, config.Password))
		log.Printf("pps:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
