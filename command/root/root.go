package root

import (
	"fmt"
	"os"

	"github/lyfeblocnetwork/lyfebloc-core/command/backup"
	"github/lyfeblocnetwork/lyfebloc-core/command/genesis"
	"github/lyfeblocnetwork/lyfebloc-core/command/helper"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft"
	"github/lyfeblocnetwork/lyfebloc-core/command/license"
	"github/lyfeblocnetwork/lyfebloc-core/command/monitor"
	"github/lyfeblocnetwork/lyfebloc-core/command/peers"
	"github/lyfeblocnetwork/lyfebloc-core/command/secrets"
	"github/lyfeblocnetwork/lyfebloc-core/command/server"
	"github/lyfeblocnetwork/lyfebloc-core/command/status"
	"github/lyfeblocnetwork/lyfebloc-core/command/txpool"
	"github/lyfeblocnetwork/lyfebloc-core/command/version"
	"github/lyfeblocnetwork/lyfebloc-core/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Lyfebloc Core is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
