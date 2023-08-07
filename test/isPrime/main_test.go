package main

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0 is not prime",
			args: args{n: 0},
			want: false,
		},
		{
			name: "1 is not prime",
			args: args{n: 1},
			want: false,
		},
		{
			name: "2 is prime",
			args: args{n: 2},
			want: true,
		},
		{
			name: "3 is prime",
			args: args{n: 3},
			want: true,
		},
		{
			name: "4 is not prime",
			args: args{n: 4},
			want: false,
		},
		{
			name: "5 is prime",
			args: args{n: 5},
			want: true,
		},
		{
			name: "-1 is not prime",
			args: args{n: -1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := isPrime(tt.args.n)
			if result != tt.want {
				t.Errorf("isPrime(%d) = %v, want %v", tt.args.n, result, tt.want)
			}
		})
	}
}
