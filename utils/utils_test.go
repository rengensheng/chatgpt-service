package utils

import "testing"

func TestTokenEncoder(t *testing.T) {
	TokenEncoder("Hello Word")
	TokenEncoder("你好世界")
	TokenEncoder("Hello你好")
}

func TestGetTokenCount(t *testing.T) {
	count := GetTokenCount("Hello Word")
	t.Log(count)
	count = GetTokenCount("你好世界")
	t.Log(count)
	count = GetTokenCount("Hello你好")
	t.Log(count)
}
