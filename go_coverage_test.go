package main

import (
	"reflect"
	"testing"
)

func Test_trimString(t *testing.T) {

	tests := []struct {
		in   string
		trim int
		want string
	}{
		{"test", 20, "test"},
		{"test", 3, "...est"},
		{"testtesttesttesttesttest", 20, "...testtesttesttesttest"},
		{"testtesttesttesttesttesttesttesttesttesttestbest", 20, "...testtesttesttestbest"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := trimString(tt.in, tt.trim); got != tt.want {
				t.Errorf("trimString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTrimmedFileName(t *testing.T) {
	type args struct {
		fn   string
		trim bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"trim enabled", args{fn: "test_file_name_test_file_name", trim: true}, "..._name_test_file_name"},
		{"trim enabled", args{fn: "test_file_name_test_file_name", trim: false}, "test_file_name_test_file_name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrimmedFileName(tt.args.fn, tt.args.trim); got != tt.want {
				t.Errorf("getTrimmedFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fmtFuncInfo(t *testing.T) {
	type args struct {
		x       *funcInfo
		tc      float64
		covered int64
		total int64
		trim    bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"returns function details without trim when trim is false",
			args{
				x: &funcInfo{
					fileName: "test",
					functionName: "test_func_name_test_func_name_test_func_name",
					functionStartLine: 10,
					functionEndLine:   20,
					uncoveredLines:    0},
				tc:      100,
				covered: 50,
				total: 50,
				trim:    false},
			[]string{"test", "test_func_name_test_func_name_test_func_name", "10", "20", "0", "0.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmtFuncInfo(tt.args.x, tt.args.tc, tt.args.covered, tt.args.total, tt.args.trim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fmtFuncInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
