package main

import (
	"testing"
)

func Test_verifyPassword(t *testing.T) {
	type args struct {
		password int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no pairs", args{123}, false},
		{"correct", args{1223}, true},
		{"not increasing", args{3223}, false},
		{"not increasing", args{3221}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPassword(tt.args.password); got != tt.want {
				t.Errorf("verifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{"OK", args{"100-200"}, 100, 200, false},
		{"hex", args{"0x100-200"}, 0, 0, true},
		{"hex", args{"100-0x200"}, 0, 0, true},
		{"characters", args{"100-a"}, 0, 0, true},
		{"characters", args{"a-100"}, 0, 0, true},
		{"not interval", args{"100"}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_verifyRange(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		// TODO: Add test cases.
		{"OK", args{246515, 739105}, false},
		{"OK", args{100000, 999999}, false},
		{"decreasing", args{739105, 246515}, true},
		{"Start not 6-digit", args{10000, 999999}, true},
		{"End not 6-digit", args{100000, 9999990}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifyRange(tt.args.start, tt.args.end); (err != nil) != tt.wantErr {
				t.Errorf("verifyRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
