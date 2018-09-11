package main

import "testing"

//代码覆盖率
//go test -coverprofile=c.out
//go tool cover -html c.out
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

//go test -bench .
//性能查看
// go test -bench . -cpuprofile cpu.out
//go tool pprof cpu.out
//web


//安装http://www.graphviz.org/ Graphviz
func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()//清空时间  删除准备数据的时间

	actual := lengthOfNonRepeatingSubStr(s)
	for i := 0; i < b.N; i++ {
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}

}
