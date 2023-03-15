package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAliAdapter_FileSystem(t *testing.T) {
	// 初始化 连接参数
	var client = AliClient{}
	// 初始化 适配器
	ali := AliAdapter{LocalClient: client}
	// 获取 操作对象
	err := ali.FileSystem()
	assert.Nil(t, err)
}

func TestAWSAdapter_FileSystem(t *testing.T) {
	var client = AWSClient{}
	aws := AWSAdapter{LocalClient: client}
	err := aws.FileSystem()
	assert.Nil(t, err)
}
