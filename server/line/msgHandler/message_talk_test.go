package msgHandler

import (
	"strings"
	"testing"
)

func Test_msgCustomized(t *testing.T) {

	tests := []string{"tc test", "tt test", "tj 測試"}
	for _, tt := range tests {
		println(processType(tt))
	}
}

func TestString(t *testing.T) {
	var a *string
	s := "TC foot"
	a = &s
	ss := *a
	if len(ss) >= 4 {
		ss = strings.ToLower(ss[:2]) + ss[2:]
	}
	println(ss)
}

func Test_msgCustomized1(t *testing.T) {
	type args struct {
		s *string
	}
	a := "TC test"
	b := "Tj 測試"
	tests := []struct {
		name   string
		args   args
		wantNs string
	}{
		// TODO: Add test cases.
		{name: "111", args: args{&b}, wantNs: "請將以下內容翻譯成中文: \"test\""},
		{"222", args{&a}, "請將以下內容翻譯成泰文: \"test\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNs := msgCustomized(tt.args.s)
			println(gotNs)

		})
	}
}
