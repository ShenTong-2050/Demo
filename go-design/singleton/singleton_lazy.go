package singleton

import "sync"

var (
	lazySingleton	*Singleton
	once			sync.Once
)

// GetLazyInstance 获取 懒汉式 单例对象
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		// 执行一次 实例化 操作
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}

