package singleton

type Singleton struct {}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance 单例获取 Singleton 对象【饿汉式】
func GetInstance() *Singleton {
	return singleton
}
