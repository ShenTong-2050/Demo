package adapter

import (
	"fmt"
)

// Flysystem 统一操作接口
type Flysystem interface {
	FileSystem()	error
}

// CreateInstance 创建连接对象接口
type CreateInstance interface {
	GetInstance()	error
}

// AWSClient AWS配置
type AWSClient struct {}

// GetInstance 连接 AWS 服务端,真实项目中应该 获取 连接后的对象
func (a *AWSClient) GetInstance() error {
	fmt.Println("connect AWS server...")
	return nil
}

// AWSAdapter AWS适配器
type AWSAdapter struct {
	LocalClient AWSClient
}

// FileSystem 实现 Flysystem 接口
func (ad *AWSAdapter) FileSystem() error {
	ad.LocalClient.GetInstance()
	return nil
}

type AliClient struct {}

func (al *AliClient) GetInstance() error {
	fmt.Println("connect to Ali Oss")
	return nil
}

type AliAdapter struct {
	LocalClient AliClient
}

func (al *AliAdapter) FileSystem() error {
	err := al.LocalClient.GetInstance()
	if err != nil {
		return err
	}
	return nil
}

