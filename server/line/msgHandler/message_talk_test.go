package msgHandler

import "testing"

func Test_msgCustomized(t *testing.T) {

	tests := []string{"tc test", "tt test", "tj 測試"}
	for _, tt := range tests {
		println(needProcess(&tt))
	}
}
