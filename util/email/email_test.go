package email

import "testing"

func TestEmail(t *testing.T) {
	SendForgotEmail("hantig1986@gmail.com", "john", "password:123456")
}
