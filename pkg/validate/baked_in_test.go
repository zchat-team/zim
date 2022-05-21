package validate

import (
	"testing"
)

func TestIsNumberGtZero(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"empty",
			"",
			false,
		},
		{
			"zero",
			"0",
			false,
		},
		{
			"多个0",
			"00",
			false,
		},
		{
			"<0",
			"-11",
			false,
		},
		{
			">0",
			"11",
			true,
		},
		{
			">0带符号",
			"+11",
			false,
		},
		{
			"浮点",
			"1.1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumberGt0(tt.args); got != tt.want {
				t.Errorf("IsNumberGt0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumberGte0(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"empty",
			"",
			false,
		},
		{
			"0",
			"0",
			true,
		},
		{
			"多个0",
			"00",
			true,
		},
		{
			"<0",
			"-11",
			false,
		},
		{
			">0",
			"11",
			true,
		},
		{
			">0带符号",
			"+11",
			false,
		},
		{
			"浮点",
			"1.1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumberGte0(tt.args); got != tt.want {
				t.Errorf("IsNumberGte0() = %v, want %v", got, tt.want)
			}
		})
	}
}
