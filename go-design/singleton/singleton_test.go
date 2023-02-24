package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试 饿汉式 单例模式 是否通过
func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

// 饿汉式 单例模式 压测
func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("testing failed!")
			}
		}
	})
}
