package adapter

import (
	"fmt"
)

type Flysystem interface {
	FileSystem()	error
}

// AWSClient AWS配置
type AWSClient struct {}

// GetInstance 连接 AWS 服务端,真实项目中应该 获取 连接后的对象
func (a *AWSClient) GetInstance() {
	fmt.Println("connect AWS server...")
}

// AWSAdapter AWS适配器
type AWSAdapter struct {
	adapter AWSClient
}

// FileSystem 实现 Flysystem 接口
func (ad *AWSAdapter) FileSystem() error {
	ad.adapter.GetInstance()
	return nil
}

type AliClient struct {}

func (al *AliClient) GetInstance() {
	fmt.Println("connect to Ali Oss")
}

type AliAdapter struct {
	adapter AliClient
}

func (al *AliAdapter) FileSystem() error {
	al.adapter.GetInstance()
	return nil
}

