package client

import (
	"log"
	"net"

	"github.com/chapin/socks5-ss/core"
)

// LsLocal struct.
type LsLocal struct {
	*core.SecureSocket
}

// New method return a LsLocal instance
func New(password *core.Password, listenAddr, remoteAddr *net.TCPAddr) *LsLocal {
	return &LsLocal{
		SecureSocket: &core.SecureSocket{
			Cipher:     core.NewCipher(password),
			ListenAddr: listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

// Listen 方法本地短启动监听，接收来自本机浏览器的连接
func (local *LsLocal) Listen(didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)
	if err != nil {
		return err
	}

	defer listener.Close()

	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		userConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// userConn 被关闭时直接清除所有的数据，不管没有发送的数据
		userConn.SetLinger(0)
		go local.handleConn(userConn)
	}

	return nil
}

// handleConn method
func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()

	proxyServer, err := local.DialRemote()
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyServer.Close()
	// Conn 被关闭时直接清除所有的数据 不管没有发送的数据
	proxyServer.SetLinger(0)

	// 进入转发
	// 从 proxyServer 读取数据发送到 localUser
	go func() {
		err := local.DecodeCopy(userConn, proxyServer)
		if err != nil {
			userConn.Close()
			proxyServer.Close()
		}
	}()

	// 从 localUser 发送数据到 proxyServer，这里因为处在翻墙阶段出现网络错误概率更大
	local.EncodeCopy(proxyServer, userConn)
}
