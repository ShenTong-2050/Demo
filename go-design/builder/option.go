package builder

import "fmt"

type ResourcePoolConfigOption struct {
	MaxTotal,MaxIdle,MinIdle int
}

type ResourcePoolConfigOptionFunc func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string,opts ...ResourcePoolConfigOptionFunc) (*ResourcePoolConfigOption,error) {
	if name == "" {
		return nil,fmt.Errorf("build args name %v cannot empty",name)
	}

	var option = &ResourcePoolConfigOption{MaxTotal: 10,MaxIdle: 9,MinIdle: 1}

	for _,opt := range opts {
		opt(option)
	}

	if option.MaxTotal == 0 || option.MaxIdle == 0 || option.MinIdle == 0 {
		return nil,fmt.Errorf("option %v args cannot be 0",option)
	}

	if option.MinIdle > option.MaxIdle || option.MaxIdle > option.MaxTotal {
		return nil,fmt.Errorf("option args %v is invalid",option)
	}

	return &ResourcePoolConfigOption{MaxTotal:option.MaxTotal,MaxIdle: option.MaxIdle,MinIdle: option.MinIdle},nil
}


