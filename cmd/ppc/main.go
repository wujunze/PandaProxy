package main

import (
	"fmt"
	"github.com/wujunze/PandaProxy/PandaProxy"
	"github.com/wujunze/PandaProxy/cmd"
	"log"
	"net"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 local 端并监听
	lsLocal, err := PandaProxy.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
远程服务地址 remote：
%s
密码 password：
%s
	`, listenAddr, config.RemoteAddr, config.Password))
		log.Printf("ppc:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
