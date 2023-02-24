package prototype

import (
	"github.com/cihub/seelog"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKeywords_Clone(t *testing.T) {
	// 获取 当前年份 2023【2006 => 为 日期模板 也就是 Go诞生的年份】
	updateTime,_ := time.Parse("2006","2023")

	var keywords = Keywords{
		"testA":&Keyword{
			word: "testA",
			visit: 1,
			updateAt: &updateTime,
		},
		"testB" : &Keyword{
			word: "testB",
			visit: 2,
			updateAt: &updateTime,
		},
		"testC" : &Keyword{
			word: "testC",
			visit: 3,
			updateAt: &updateTime,
		},
	}

	now := time.Now()

	updateKeywords := []*Keyword{
		{
			word: "testB",
			visit: 10,
			updateAt: &now,
		},
	}

	got := keywords.Clone(updateKeywords)

	log.Debugf("time now => %v",now)
	seelog.Info(keywords["testA"],got["testA"])
	seelog.Info(keywords["testB"],got["testB"])

	assert.Equal(t, keywords["testA"],got["testA"])				// PASS
	assert.NotEqual(t, keywords["testB"],got["testB"])			// keywords["testB"] 与 got["testB"] 不相等 所以 测试 PASS
	assert.NotEqual(t, keywords["testB"],updateKeywords[0])		// keywords["testB"] 与 updateKeywords[0] 不相等 所以 测试 PASS
	assert.Equal(t, keywords["testC"],got["testC"])				// keywords["testC"] 与 got["testC"] 相等 所以 测试 PASS

	// 最后一个测试用例不通过
	// require.Equalf(t, keywords["testB"],got["testB"],"keywords testB %v not equal got testB %v",keywords["testB"],got["testB"])
}
