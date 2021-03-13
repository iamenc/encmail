package mail

import "github.com/go-gomail/gomail"

// Test 测试
func sum(a ...int) (r int) {
	for _, i := range a {
		r += i
	}
	return
}

// OK 邮件实例
func OK() {
	d := gomail.NewDialer()
}
