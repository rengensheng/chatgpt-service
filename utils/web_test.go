package utils

import "testing"

func TestGetWebText(t *testing.T) {
	text, err := GetWebText("https://vitejs.dev/guide/build.html")
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(text)
	}
}
