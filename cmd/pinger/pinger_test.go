package pinger_test

import (
	"testing"

	"github.com/Nayls/pinger/cmd/completion"
	"github.com/Nayls/pinger/cmd/generate"
	"github.com/Nayls/pinger/cmd/server"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/Nayls/pinger/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func buildTestCmd() *cobra.Command {
	cmd := cli.GetRootCmd()

	// Add generate command
	cmd.AddCommand(generate.GetGenerateCmd())

	// Add server command
	cmd.AddCommand(server.GetServerCmd())

	// Add completion command
	cmd.AddCommand(completion.GetCompletionCmd())

	return cmd
}

func Test_SingleCommandWithoutSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Application for pinger another host`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown command "fail" for "pinger"`)

	assert.NotContains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.NotContains(t, out, `Flags:`)

	assert.Contains(t, out, `Run 'pinger --help' for usage.`)
}

func Test_NegativeCallSingleCommandWithLongFailFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown flag: --fail`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}
