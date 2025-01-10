package peers

import (
	"github/lyfeblocnetwork/lyfebloc-core/command/helper"
	"github/lyfeblocnetwork/lyfebloc-core/command/peers/add"
	"github/lyfeblocnetwork/lyfebloc-core/command/peers/list"
	"github/lyfeblocnetwork/lyfebloc-core/command/peers/status"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	peersCmd := &cobra.Command{
		Use:   "peers",
		Short: "Top level command for interacting with the network peers. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(peersCmd)

	registerSubcommands(peersCmd)

	return peersCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// peers status
		status.GetCommand(),
		// peers list
		list.GetCommand(),
		// peers add
		add.GetCommand(),
	)
}
