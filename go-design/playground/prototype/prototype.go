package prototype

import (
	"encoding/json"
)

type UserData struct {
	name string
	age  int
}

func (u UserData) Clone() *UserData {
	// 深拷贝 用户 数据
	var newUserData UserData
	b,_ := json.Marshal(u)
	json.Unmarshal(b,newUserData)
	return &newUserData
}

type UserDatas map[string]*UserData

func (us UserDatas) Clones(updateUserData []*UserData) UserDatas {
	var tmpUserDatas = UserDatas{}
	for k,v := range us {
		tmpUserDatas[k] = v
	}
	for _,u := range updateUserData {
		// 浅拷贝
		tmpUserDatas[u.name] = u.Clone()
	}
	return tmpUserDatas
}
