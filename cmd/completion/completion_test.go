package completion_test

import (
	"testing"

	"github.com/Nayls/pinger/cmd/completion"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/Nayls/pinger/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func buildTestCmd() *cobra.Command {
	cmd := cli.GetRootCmd()

	// Add completion command
	cmd.AddCommand(completion.GetCompletionCmd())

	return cmd
}

func Test_SingleCommandWithoutSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "completion")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: accepts 1 arg(s), received 0`)
	assert.Contains(t, out, `completion [bash|zsh|fish|powershell]`)
}

func Test_SingleCommandWithLongFlagHelp(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "completion", "--help")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)

	assert.Contains(t, out, `To load completions:`)
	assert.Contains(t, out, `Bash:`)
	assert.Contains(t, out, `Zsh:`)
	assert.Contains(t, out, `Fish:`)
	assert.Contains(t, out, `PowerShell:`)
}

func Test_SingleCommandWithShortFlagHelp(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "completion", "-h")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)

	assert.Contains(t, out, `To load completions:`)
	assert.Contains(t, out, `Bash:`)
	assert.Contains(t, out, `Zsh:`)
	assert.Contains(t, out, `Fish:`)
	assert.Contains(t, out, `PowerShell:`)
}

func Test_NegativeCallSingleCommandWithSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "completion", "fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error: invalid argument "fail" for "pinger completion"`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithoutSubcommandWithFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "completion", "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, "Error: unknown flag: --fail")

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}
