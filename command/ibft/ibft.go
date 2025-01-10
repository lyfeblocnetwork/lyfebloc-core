package ibft

import (
	"github/lyfeblocnetwork/lyfebloc-core/command/helper"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft/candidates"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft/propose"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft/quorum"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft/snapshot"
	"github/lyfeblocnetwork/lyfebloc-core/command/ibft/status"
	_switch "github/lyfeblocnetwork/lyfebloc-core/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
