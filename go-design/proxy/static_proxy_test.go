package proxy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProxy_Login(t *testing.T) {
	// 获取 proxyUser 结构体
	proxyUser := NewUserProxy(&User{})
	// 通过代理 实现登录 流程
	err := proxyUser.Login("username","password")
	// 判断登录结果是否异常
	require.Nil(t, err)
}
