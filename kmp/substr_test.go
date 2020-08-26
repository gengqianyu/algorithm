package kmp

import (
	"testing"
)

func TestMaxLenSubstr(t *testing.T) {
	tests := []struct {
		s   string
		len int
		str string
	}{
		//edge case
		{"abc", 3, "abc"},
		{"abcfabcbacbfdcada", 5, "acbfd"},
		// chinese support
		{"耿洋洋", 2, "耿洋"},
	}

	for _, tt := range tests {
		if l, ms := MaxLenSubstr(tt.s); l != tt.len || ms != tt.str {
			t.Errorf("MaxLenSubstr(%s),got maxlen：%d, maxstr:%s;expected maxlen：%d, maxstr:%s", tt.s, l, ms, tt.len, tt.str)
		}
	}
}

type StringIndexFunc func(string, string) int

func TestIndexStrings(t *testing.T) {
	marshalTests := []struct {
		name   string
		new    func() StringIndexFunc
		str    string
		substr string
		index  int
	}{
		{
			"stringsIndex",
			func() StringIndexFunc { return IndexStrings },
			"gengqianyuzhangyanlin",
			"zhangyanlin",
			10,
		},
		{
			"stringsIndex",
			func() StringIndexFunc { return IndexStrings },
			"白化肥挥发花化肥黑化肥挥发白化肥",
			"花化肥黑",
			5,
		},
	}

	for _, tt := range marshalTests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.new
			i := f()(tt.str, tt.substr)
			if i != tt.index {
				t.Errorf("expected:%d,actual:%d", tt.index, i)
			}
		})
	}
}
