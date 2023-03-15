package adapter

import "testing"

func TestAliAdapter_FileSystem(t *testing.T) {
	// 初始化 连接参数
	var client = AliClient{}
	// 初始化 适配器
	aliOss := AliAdapter{adapter: client}
	// 获取 操作对象
	aliOss.FileSystem()
}

func TestAWSAdapter_FileSystem(t *testing.T) {
	var client = AWSClient{}
	aws := AWSAdapter{adapter: client}
	aws.FileSystem()
}
