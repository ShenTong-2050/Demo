package facade

type User struct {
	name string
}

// IUser 用户登入所有 行为
type IUser interface {
	Login(phone,code int) (*User,error)
	Register(phone,code int) (*User,error)
}

// IUserFacade 需要整合的操作
type IUserFacade interface {
	LoginOrRegister(phone,code int) (*User,error)
}

// UserService 需要整合的 行为
type UserService struct {}

func (u UserService) Login(phone,code int) (*User,error) {
	// 校验 登录
	// return &User{name: "test login..."},nil
	return nil,nil
}

func (u UserService) Register(phone,code int) (*User,error) {
	// 校验 注册
	return &User{name: "test register..."},nil
}

func (u UserService) LoginOrRegister(phone,code int) (*User,error) {

	user,err := u.Login(phone,code)

	// 校验失败 或 用户不存在
	if err != nil {
		return nil,err
	}

	if user != nil {
		return user,nil
	}

	// 注册
	user,err = u.Register(phone,code)

	return user,nil
}

