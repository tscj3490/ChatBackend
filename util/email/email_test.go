package email

import "testing"

func TestEmail(t *testing.T) {
	SendForgotEmail("hantig1986@gmail.com", "password:123456")
}
