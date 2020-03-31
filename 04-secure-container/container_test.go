package main

import "testing"

func Test_verifyPassword(t *testing.T) {
	type args struct {
		password int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
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
