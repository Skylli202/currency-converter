package cmd_test

import (
	"testing"

	"github.com/Skylli202/currency-converter/cmd"
	"github.com/spf13/cobra"
)

func TestValidateConvertCmdArgs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cmd     *cobra.Command
		args    []string
		wantErr bool
	}{
		{
			name:    "no args",
			cmd:     nil,
			args:    []string{},
			wantErr: false,
		},
		{
			name:    "one arg - valid",
			cmd:     nil,
			args:    []string{"123"},
			wantErr: false,
		},
		{
			name:    "one arg - invalid (NaN)",
			cmd:     nil,
			args:    []string{"abc"},
			wantErr: true,
		},
		{
			name:    "one arg - invalid (empty)",
			cmd:     nil,
			args:    []string{""},
			wantErr: true,
		},
		{
			name:    "two arg - valid",
			cmd:     nil,
			args:    []string{"123", "CAD"},
			wantErr: false,
		},
		{
			name:    "two arg - valid (reverse order)",
			cmd:     nil,
			args:    []string{"CAD", "123"},
			wantErr: false,
		},
		{
			name:    "two arg - invalid (no amount)",
			cmd:     nil,
			args:    []string{"foo", "bar"},
			wantErr: true,
		},
		{
			name:    "two arg - invalid (not a valid currency code)",
			cmd:     nil,
			args:    []string{"123", "NOT A VALID CURRENCY CODE"},
			wantErr: true,
		},
		{
			name:    "two arg - invalid (not a valid currency code reverse)",
			cmd:     nil,
			args:    []string{"NOT A VALID CURRENCY CODE", "123"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := cmd.ValidateConvertCmdArgs(tt.cmd, tt.args)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ValidateConvertCmdArgs() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ValidateConvertCmdArgs() succeeded unexpectedly")
			}
		})
	}
}
