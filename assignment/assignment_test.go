package assignment

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAddUint32(t *testing.T) {
	type args struct {
		x uint32
		y uint32
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{"case1", args{math.MaxUint32, 1}, 0, true},
		{"case2", args{1, 1}, 2, false},
		{"case3", args{42, 2701}, 2743, false},
		{"case4", args{42, math.MaxUint32}, 41, true},
		{"case5", args{4294967290, 5}, 4294967295, false},
		{"case6", args{4294967290, 6}, 0, true},
		{"case7", args{4294967290, 10}, 4, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, overflow := AddUint32(tt.args.x, tt.args.y)
			assert.Equalf(t, tt.want, sum, "AddUint32(%v, %v)", tt.args.x, tt.args.y)
			assert.Equalf(t, tt.want1, overflow, "AddUint32(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}

func TestCeilNumber(t *testing.T) {
	tests := []struct {
		name string
		f    float64
		want float64
	}{
		{"case1", 42.42, 42.50},
		{"case2", 42, 42},
		{"case3", 42.01, 42.25},
		{"case4", 42.24, 42.25},
		{"case5", 42.25, 42.25},
		{"case6", 42.26, 42.50},
		{"case7", 42.55, 42.75},
		{"case8", 42.75, 42.75},
		{"case9", 42.76, 43},
		{"case10", 42.99, 43},
		{"case11", 43.13, 43.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CeilNumber(tt.f), "CeilNumber(%v)", tt.f)
		})
	}
}

func TestAlphabetSoup(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"case1", "hello", "ehllo"},
		{"case2", "", ""},
		{"case3", "h", "h"},
		{"case4", "ab", "ab"},
		{"case5", "ba", "ab"},
		{"case6", "bac", "abc"},
		{"case7", "cba", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AlphabetSoup(tt.s), "AlphabetSoup(%v)", tt.s)
		})
	}
}

func TestStringMask(t *testing.T) {
	type args struct {
		s string
		n uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"!mysecret*", 2}, "!m********"},
		{"case2", args{"", 1}, "*"},
		{"case3", args{"a", 1}, "*"},
		{"case4", args{"string", 0}, "******"},
		{"case5", args{"string", 3}, "str***"},
		{"case6", args{"string", 5}, "strin*"},
		{"case7", args{"string", 6}, "******"},
		{"case8", args{"string", 7}, "******"},
		{"case9", args{"s*r*n*", 3}, "s*r***"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StringMask(tt.args.s, tt.args.n), "StringMask(%v, %v)", tt.args.s, tt.args.n)
		})
	}
}

func TestWordSplit(t *testing.T) {
	words := "apple,bat,cat,goodbye,hello,yellow,why"
	tests := []struct {
		name string
		arr  [2]string
		want string
	}{
		{"case1", [2]string{"hellocat", words}, "hello, cat"},
		{"case2", [2]string{"catbat", words}, "cat, bat"},
		{"case3", [2]string{"yellowapple", words}, "yellow, apple"},
		{"case4", [2]string{"", words}, "not possible"},
		{"case5", [2]string{"notcat", words}, "not possible"},
		{"case6", [2]string{"bootcamprocks!", words}, "not possible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WordSplit(tt.arr), "WordSplit(%v)", tt.arr)
		})
	}
}

func TestVariadicSet(t *testing.T) {
	tests := []struct {
		name string
		i    []interface{}
		want []interface{}
	}{
		{"case1", []interface{}{4, 2, 5, 4, 2, 4}, []interface{}{4, 2, 5}},
		{"case2", []interface{}{"bootcamp", "rocks!", "really", "rocks!"}, []interface{}{"bootcamp", "rocks!", "really"}},
		{"case3", []interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"}, []interface{}{1, uint32(1), "first", 2, uint32(2), "second"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, VariadicSet(tt.i...), "VariadicSet(%v)", tt.i...)
		})
	}
}

func BenchmarkStringMask(b *testing.B) {
	type args struct {
		s string
		n uint
	}
	benchmarks := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"!mysecret*", 2}, "!m********"},
		{"case2", args{"", 1}, "*"},
		{"case3", args{"a", 1}, "*"},
		{"case4", args{"string", 0}, "******"},
		{"case5", args{"string", 3}, "str***"},
		{"case6", args{"string", 5}, "strin*"},
		{"case7", args{"string", 6}, "******"},
		{"case8", args{"string", 7}, "******"},
		{"case9", args{"s*r*n*", 3}, "s*r***"},
	}

	for _, v := range benchmarks {
		b.Run(fmt.Sprintf("inputsize%T", v.args.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringMask(v.args.s, v.args.n)
			}
		})
	}
}

func BenchmarkStringMaskReyyanSolve(b *testing.B) {
	type args struct {
		s string
		n uint
	}
	benchmarks := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"!mysecret*", 2}, "!m********"},
		{"case2", args{"", 1}, "*"},
		{"case3", args{"a", 1}, "*"},
		{"case4", args{"string", 0}, "******"},
		{"case5", args{"string", 3}, "str***"},
		{"case6", args{"string", 5}, "strin*"},
		{"case7", args{"string", 6}, "******"},
		{"case8", args{"string", 7}, "******"},
		{"case9", args{"s*r*n*", 3}, "s*r***"},
	}

	for _, v := range benchmarks {
		b.Run(fmt.Sprintf("inputsize%T", v.args.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringMaskReyyanSolve(v.args.s, v.args.n)
			}
		})
	}
}
