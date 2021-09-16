package generate_test

import (
	"os"
	"testing"

	"github.com/Nayls/pinger/cmd/generate"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/Nayls/pinger/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func buildTestCmd() *cobra.Command {
	cmd := cli.GetRootCmd()

	// Add generate command
	cmd.AddCommand(generate.GetGenerateCmd())

	return cmd
}

func Test_CallSingleCommandWithoutSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "generate")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Command for generate docs and etc`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithSubcommand(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "generate", "fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Command for generate docs and etc`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithoutSubcommandWithFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "generate", "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown flag: --fail`)

	assert.Contains(t, out, `Usage:`)
	assert.Contains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)

}

func Test_CallCommandWithSubcommandCli(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "generate", "cli")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)
	assert.Empty(t, out)
	assert.DirExists(t, "./docs/cli")
	assert.FileExists(t, "./docs/cli/pinger.md")
	assert.FileExists(t, "./docs/cli/pinger_generate.md")
	assert.FileExists(t, "./docs/cli/pinger_generate_cli.md")

	os.RemoveAll("./docs")
}

func Test_NegativeCallCommandWithSubcommandCliWithFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "generate", "cli", "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, `Error: unknown flag: --fail`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}
