package configs

import (
	"fmt"
	"testing"
)

func TestNewConfigs(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestNewConfigs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := NewConfigs("../")
			fmt.Println(config)
		})
	}
}
