package main

import (
	"testing"
)

func Test_isAngka(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "angka ",
			args: args{param: "1"},
			want: true},
		{name: "bukan angka ",
			args: args{param: "a"},
			want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAngka(tt.args.param); got != tt.want {
				t.Errorf("isAngka() = %v, want %v", got, tt.want)
			}
		})
	}
}
