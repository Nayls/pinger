package server_test

import (
	"testing"

	"github.com/Nayls/pinger/cmd/server"
	"github.com/Nayls/pinger/internal/cli"
	"github.com/Nayls/pinger/internal/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func buildTestCmd() *cobra.Command {
	cmd := cli.GetRootCmd()

	// Add server command
	cmd.AddCommand(server.GetServerCmd())

	return cmd
}

func Test_SingleCommandWithLongFlagHelp(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "server", "--help")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_SingleCommandWithShortFlagHelp(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "server", "-h")
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotContains(t, out, `Error:`)

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}

func Test_NegativeCallSingleCommandWithLongFailFlag(t *testing.T) {
	out, err := test.ExecuteCommand(buildTestCmd(), "server", "--fail")
	if err != nil {
		assert.Error(t, err)
	}

	assert.Contains(t, out, "Error: unknown flag: --fail")

	assert.Contains(t, out, `Usage:`)
	assert.NotContains(t, out, `Available Commands:`)
	assert.Contains(t, out, `Flags:`)
}
