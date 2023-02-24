package singleton

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

// 测试懒汉单例模式是否通过
func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

// 懒汉单例模式压测
func BenchmarkGetLazyParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("testing failed!")
			}
		}
	})
}

func BenchmarkTemplateParallel(b *testing.B) {
	temp := template.Must(template.New("test").Parse("Hello {{.}} !"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			temp.Execute(&buf,"world")
		}
	})
}
