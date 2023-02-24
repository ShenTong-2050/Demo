package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word string
	visit int
	updateAt *time.Time
}

type Keywords map[string]*Keyword

func (word *Keyword) Clone() *Keyword {
	var newKeyWord Keyword
	b,_ := json.Marshal(word)		// 使用 序列化 与 反序列化 的方式 深拷贝
	json.Unmarshal(b,newKeyWord)
	return &newKeyWord
}

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	var newKeyWords = Keywords{}

	// 记录 临时变量
	for k,v := range words {
		newKeyWords[k] = v 		// 这里是 浅拷贝 直接拷贝了 地址
	}

	for _,word := range updateWords {
		newKeyWords[word.word] = word.Clone()	// 替换掉 需要拷贝的地址 这里是 深拷贝
	}
	return newKeyWords
}
