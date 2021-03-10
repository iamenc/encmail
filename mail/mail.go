package mail

// Test 测试
func sum(a ...int) (r int) {
	for _, i := range a {
		r += i
	}
	return
}
