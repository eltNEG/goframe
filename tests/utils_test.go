package tests

import (
	"fmt"
	"goframe/utils"
	"testing"
)

func TestGenerateID(t *testing.T) {
	tests := []struct {
		name       string
		prefix     string
		surfix     string
		wantPrefix string
		wantSurfix string
	}{
		{
			name:       "TestGenerateID with no prefix and surfix",
			prefix:     "",
			surfix:     "",
			wantPrefix: "",
			wantSurfix: "",
		},
		{
			name:       "TestGenerateID with prefix and no surfix",
			prefix:     "prefix",
			surfix:     "",
			wantPrefix: "prefix",
			wantSurfix: "",
		},
		{
			name:       "TestGenerateID with no prefix and surfix",
			prefix:     "",
			surfix:     "surfix",
			wantPrefix: "",
			wantSurfix: "surfix",
		},
		{
			name:       "TestGenerateID with prefix and surfix",
			prefix:     "prefix",
			surfix:     "surfix",
			wantPrefix: "prefix",
			wantSurfix: "surfix",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GenID(utils.GenIDArg.WithPrefix(tt.prefix), utils.GenIDArg.WithSurfix(tt.surfix))
			fmt.Println(got)
			if tt.wantPrefix != "" {
				if got[:len(tt.wantPrefix)] != tt.wantPrefix {
					t.Errorf("GenerateID() = %v, want %v", got, tt.wantPrefix)
				}
			}
			if tt.wantSurfix != "" {
				if got[len(got)-len(tt.wantSurfix):] != tt.wantSurfix {
					t.Errorf("GenerateID() = %v, want %v", got, tt.wantSurfix)
				}
			}
		})
	}
}
