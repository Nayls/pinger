package test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func EmptyRun(*cobra.Command, []string) {}

func ExecuteCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = ExecuteCommandC(root, args...)
	return output, err
}

func ExecuteCommandWithContext(ctx context.Context, root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err = root.ExecuteContext(ctx)

	return buf.String(), err
}

func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func ExecuteCommandWithContextC(ctx context.Context, root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteContextC(ctx)

	return c, buf.String(), err
}

func ResetCommandLineFlagSet() {
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
}

func CheckStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}

func CheckStringOmits(t *testing.T, got, expected string) {
	if strings.Contains(got, expected) {
		t.Errorf("Expected to not contain: \n %v\nGot: %v", expected, got)
	}
}
