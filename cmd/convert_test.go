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
			name:    "two args - valid",
			cmd:     nil,
			args:    []string{"123", "CAD"},
			wantErr: false,
		},
		{
			name:    "two args - valid (reverse order)",
			cmd:     nil,
			args:    []string{"CAD", "123"},
			wantErr: false,
		},
		{
			name:    "two args - invalid (no amount)",
			cmd:     nil,
			args:    []string{"foo", "bar"},
			wantErr: true,
		},
		{
			name:    "two args - invalid (not a valid currency code)",
			cmd:     nil,
			args:    []string{"123", "NOT A VALID CURRENCY CODE"},
			wantErr: true,
		},
		{
			name:    "two args - invalid (not a valid currency code reverse)",
			cmd:     nil,
			args:    []string{"NOT A VALID CURRENCY CODE", "123"},
			wantErr: true,
		},
		{
			name:    "three args - valid",
			cmd:     nil,
			args:    []string{"1234", "EUR", "CAD"},
			wantErr: false,
		},
		{
			name:    "three args - valid (reverse order)",
			cmd:     nil,
			args:    []string{"EUR", "CAD", "1234"},
			wantErr: false,
		},
		{
			name:    "three args - invalid (invalid order)",
			cmd:     nil,
			args:    []string{"EUR", "1234", "CAD"},
			wantErr: true,
		},
		{
			name:    "three args - invalid (invalid currency code)",
			cmd:     nil,
			args:    []string{"EUR", "NOT A VALID CURRENCY CODE", "1234"},
			wantErr: true,
		},
		{
			name:    "three args - invalid (invalid currency code)",
			cmd:     nil,
			args:    []string{"NOT A VALID CURRENCY CODE", "EUR", "1234"},
			wantErr: true,
		},
		{
			name:    "three args - invalid (invalid currency code)",
			cmd:     nil,
			args:    []string{"NOT VALID", "ALSO NOT VALID", "1234"},
			wantErr: true,
		},
		{
			name:    "three args - invalid (invalid currency code)",
			cmd:     nil,
			args:    []string{"1234", "NOT VALID", "ALSO NOT VALID"},
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
