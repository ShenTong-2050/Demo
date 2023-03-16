package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	phone 	int = 13111111111
	code 	int = 123456
)

func TestUserService_Login(t *testing.T) {
	usl,err := UserService{}.Login(phone,code)
	assert.NoError(t, err)
	assert.Equal(t, &User{name: "test login..."},usl)
}

func TestUserService_Register(t *testing.T) {
	usr,err := UserService{}.Register(phone,code)
	assert.NoError(t, err)
	assert.Equal(t, usr,&User{name: "test register..."})
}

func TestUserService_LoginOrRegister(t *testing.T) {
	uslr,err := UserService{}.LoginOrRegister(phone,code)
	assert.NoError(t, err)
	assert.Equal(t, uslr,&User{name: "test register..."})
}
