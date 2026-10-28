package entiers

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []int
		wantErr bool
	}{
		{
			name: "separate arguments",
			args: []string{"3", "2", "1", "0"},
			want: []int{3, 2, 1, 0},
		},
		{
			name: "space-separated argument",
			args: []string{"3 2 1 0"},
			want: []int{3, 2, 1, 0},
		},
		{
			name: "empty arguments",
			args: []string{},
			want: []int{},
		},
		{
			name:    "whitespace-only argument",
			args:    []string{"   "},
			wantErr: true,
		},
		{
			name:    "empty argument between values",
			args:    []string{"3", "", "2"},
			wantErr: true,
		},
		{
			name:    "non-integer",
			args:    []string{"3", "abc"},
			wantErr: true,
		},
		{
			name:    "duplicate",
			args:    []string{"3", "2", "3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Parse(%v) error = %v, wantErr %t", tt.args, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !equalInts(got, tt.want) {
				t.Errorf("Parse(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func equalInts(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
