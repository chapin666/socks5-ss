package main

import (
	"fmt"
	"log"
	"net"

	"github.com/chapin/socks5-ss/client"
	"github.com/chapin/socks5-ss/core"

	"github.com/chapin/socks5-ss/cmd"
)

const (
	// DefaultListenAddr .
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

	// 解析配置
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动local端并监听
	lsLocal := client.New(password, listenAddr, remoteAddr)
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("使用配置：", fmt.Sprintf(`
			本地监听地址 listen：
			%s
			远程服务地址 remote：
			%s
			密码 password：
			%s
		 `), listenAddr, remoteAddr, password)
		log.Printf("lightsocks-local:%s 启动成功 监听在 %s\n", version, listenAddr.String())
	}))
}
