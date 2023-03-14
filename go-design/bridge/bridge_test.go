package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewErrorNotification(t *testing.T) {
	// 获取 发送邮件行为
	emailSender := NewEmailMsgSender([]string{"test@qq.com"})
	emailNotify := NewErrorNotification(emailSender)
	err := emailNotify.Notify("send content....")
	assert.Nil(t, err)
}
