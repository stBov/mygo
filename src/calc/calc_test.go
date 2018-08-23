package src_calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if res := Add(6, 2); res != 8 {
		t.Error("加法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}