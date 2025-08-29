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

func TestParseConvertCmdArgs(t *testing.T) {
	// Currency code test cases are commented because the current implementation
	// of ParseconvertCmdArgs does not return an error. Plus, this already has
	// been validated by the validation of argument.
	tests := []struct {
		name           string
		cmd            *cobra.Command
		args           []string
		wantAmt        float64
		wantFromCurr   string
		wantTargetCurr string
	}{
		{
			name:           "no args",
			cmd:            nil,
			args:           []string{},
			wantAmt:        0,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		{
			name:           "one arg - valid",
			cmd:            nil,
			args:           []string{"123.45"},
			wantAmt:        123.45,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		{
			name:           "one arg - valid",
			cmd:            nil,
			args:           []string{"54.123"},
			wantAmt:        54.123,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		{
			name:           "one arg - invalid (NaN)",
			cmd:            nil,
			args:           []string{"abc"},
			wantAmt:        0,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		{
			name:           "one arg - invalid (empty)",
			cmd:            nil,
			args:           []string{""},
			wantAmt:        0,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		{
			name:           "two args - valid",
			cmd:            nil,
			args:           []string{"123.45", "CAD"},
			wantAmt:        123.45,
			wantFromCurr:   "USD",
			wantTargetCurr: "CAD",
		},
		{
			name:           "two args - valid (reverse order)",
			cmd:            nil,
			args:           []string{"EUR", "123.45"},
			wantAmt:        123.45,
			wantFromCurr:   "USD",
			wantTargetCurr: "EUR",
		},
		{
			name:           "two args - invalid (no amount)",
			cmd:            nil,
			args:           []string{"foo", "bar"},
			wantAmt:        0,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		//{
		//	name:           "two args - invalid (not a valid currency code)",
		//	cmd:            nil,
		//	args:           []string{"123", "NOT A VALID CURRENCY CODE"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
		//{
		//	name:           "two args - invalid (not a valid currency code reverse)",
		//	cmd:            nil,
		//	args:           []string{"NOT A VALID CURRENCY CODE", "123"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
		{
			name:           "three args - valid",
			cmd:            nil,
			args:           []string{"1234", "EUR", "CAD"},
			wantAmt:        1234,
			wantFromCurr:   "EUR",
			wantTargetCurr: "CAD",
		},
		{
			name:           "three args - valid (reverse order)",
			cmd:            nil,
			args:           []string{"GBP", "CAD", "5678"},
			wantAmt:        5678,
			wantFromCurr:   "GBP",
			wantTargetCurr: "CAD",
		},
		{
			name:           "three args - invalid (invalid order)",
			cmd:            nil,
			args:           []string{"EUR", "1234", "CAD"},
			wantAmt:        0,
			wantFromCurr:   "",
			wantTargetCurr: "",
		},
		//{
		//	name:           "three args - invalid (invalid currency code)",
		//	cmd:            nil,
		//	args:           []string{"EUR", "NOT A VALID CURRENCY CODE", "1234"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
		//{
		//	name:           "three args - invalid (invalid currency code)",
		//	cmd:            nil,
		//	args:           []string{"NOT A VALID CURRENCY CODE", "EUR", "1234"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
		//{
		//	name:           "three args - invalid (invalid currency code)",
		//	cmd:            nil,
		//	args:           []string{"NOT VALID", "ALSO NOT VALID", "1234"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
		//{
		//	name:           "three args - invalid (invalid currency code)",
		//	cmd:            nil,
		//	args:           []string{"1234", "NOT VALID", "ALSO NOT VALID"},
		//	wantAmt:        0,
		//	wantFromCurr:   "",
		//	wantTargetCurr: "",
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amt, fromCurr, targetCurr := cmd.ParseConvertCmdArgs(tt.cmd, tt.args)

			if tt.wantAmt != amt || tt.wantFromCurr != fromCurr || tt.wantTargetCurr != targetCurr {
				t.Errorf("ParseConvertCmdArgs(%v) = %#v %#v %#v, wanted %#v %#v %#v", tt.args, amt, fromCurr, targetCurr,
					tt.wantAmt,
					tt.wantFromCurr, tt.wantTargetCurr)
			}
		})
	}
}
