package builder

import (
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle 	= 9
	defaultMinIdle  = 1
)

type resourcePoolConfig struct {
	name   		string
	maxTotal	int
	maxIdle		int
	minIdle     int
}

type resourcePoolConfigBuilder struct {
	name 		string
	maxTotal	int
	maxIdle		int
	minIdle		int
}

func (b *resourcePoolConfigBuilder) setName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot empty %v",name)
	}
	b.name = name
	return nil
}

func (b *resourcePoolConfigBuilder) setMaxTotal(maxTotal int) error {
	if maxTotal < 0 {
		return fmt.Errorf("maxTotal %v cannot < 0",maxTotal)
	}
	b.maxTotal = maxTotal
	return nil
}

func (b *resourcePoolConfigBuilder) setMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("maxIdle %v cannot < 0",maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *resourcePoolConfigBuilder) setMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("minIdle %v cannot < 0",minIdle)
	}
	b.minIdle = minIdle
	return nil
}

func (b *resourcePoolConfigBuilder) Build() (*resourcePoolConfig,error) {
	if b.name == "" {
		return nil,fmt.Errorf("resource pool name %v not empty",b.name)
	}

	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}

	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}

	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}

	if b.maxIdle > b.maxTotal {
		return nil,fmt.Errorf("resource pool maxIdle %v can not > maxTotal %v",b.maxIdle,b.maxTotal)
	}

	if b.minIdle > b.maxIdle {
		return nil,fmt.Errorf("resource pool minIdle %v can not > maxIdle %v",b.minIdle,b.maxIdle)
	}

	return &resourcePoolConfig {
		name: b.name,
		maxTotal: b.maxTotal,
		maxIdle: b.maxIdle,
		minIdle: b.minIdle,
	} , nil
}













