package mail

import "testing"

// TestSum 测试单元
func TestSum(t *testing.T) {
	if sum(1, 2, 3, 4, 5, 6, 7, 8) != 36 {
		t.Error("chuowu")
	}
}
