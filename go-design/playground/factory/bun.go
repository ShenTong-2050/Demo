package factory

/**
 * 工厂方法 实现步骤：
 * 1. 一个 返回参数 为 产品接口 的 工厂接口
 * 2. 一个产品接口
 * 3. 细分产品 结构体 与 接口的实现【上海肉馅包子、上海素馅包子、北京肉馅包子、北京素馅包子】
 * 4. 细分工厂 结构体 与 接口的实现【上海的工厂、北京的工厂】
 */

// Buns 产品接口
type Buns interface {
	create() string
}

// BunShop 工厂接口
type BunShop interface {
	Generate(t string) Buns
}

// ----------------------------- 产品 与 工厂 接口 定义完成

// BJMeatBuns 北京肉馅结构体
type BJMeatBuns struct {}

// create BJMeatBuns 结构体 实现 Buns 接口
func (*BJMeatBuns) create() string {
	return "Beijing meat buns"
	// fmt.Println("Beijing meat buns")
}

// BJVegetableBuns 北京素馅结构体
type BJVegetableBuns struct {}

// create BJVegetableBuns 结构体 实现 Buns 接口
func (*BJVegetableBuns) create() string {
	return "Beijing vegetable buns"
	// fmt.Println("Beijing vegetable buns")
}

// SHMeatBuns 上海肉馅结构体
type SHMeatBuns struct {}

// create SHMeatBuns 结构体 实现 Buns 接口
func (*SHMeatBuns) create() string {
	return "Shanghai meat buns"
	// fmt.Print("Shanghai meat buns")
}

// SHVegetable 上海素馅接口
type SHVegetable struct {}

// create SHVegetable 结构体 实现 Buns 接口
func (*SHVegetable) create() string {
	return "Shanghai vegetable buns"
	// fmt.Println("Shanghai vegetable buns")
}

// ------------------ 至此 产品 细分完成

// SHFactory 上海工厂结构体
type SHFactory struct {}

// BJFactory 北京工厂结构体
type BJFactory struct {}

// Generate SHFactory 工厂实现 BunShop 接口
func (*SHFactory) Generate(t string) Buns {
	switch t {
	case "meat":
		return new(SHMeatBuns)
	case "vegetable":
		return new(SHVegetable)
	}
	return nil
}

// Generate BJFactory 工厂实现 BunShop 接口
func (*BJFactory) Generate(t string) Buns {
	switch t {
	case "meat":
		return new(BJMeatBuns)
	case "vegetable":
		return new(BJVegetableBuns)
	}
	return nil
}

// --------------------------- 至此 工厂 细分 完成
