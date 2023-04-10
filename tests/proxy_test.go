package test

import (
	"testing"

	"github.com/code-scan/Goal/Gproxy"
)

func TestProxy(t *testing.T) {
	go Gproxy.ClientWait("8888") //监听客户端端口，等待客户端链接
	go Gproxy.ServerWait("8889") // 服务端端口，用户链接用于代理
	Gproxy.RunProxy("127.0.0.1:8888")
	// curl baidu.com -x socks5://127.0.0.1:8889
}
