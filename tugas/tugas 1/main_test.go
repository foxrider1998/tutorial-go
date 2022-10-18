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
		{name: "succes",
			args: args{param: "1"},
			want: true},
		{name: "failed",
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

func Test_isVocal(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "succes",
			args: args{param: "a"},
			want: true},
		{name: "failed",
			args: args{param: "p"},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVocal(tt.args.param); got != tt.want {
				t.Errorf("isVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iskata(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "succes",
			args: args{param: "1"},
			want: false},
		{name: "failed",
			args: args{param: "aku"},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iskata(tt.args.param); got != tt.want {
				t.Errorf("iskata() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ispolindrone(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "succes",
			args: args{input: "madam"},
			want: true},
		{name: "failed",
			args: args{input: "mama"},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ispolindrone(tt.args.input); got != tt.want {
				t.Errorf("ispolindrone() = %v, want %v", got, tt.want)
			}
		})
	}
}
