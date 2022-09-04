package root

import (
	"fmt"
	"os"

	"github.com/JResearchLabs/Flutechain/command/backup"
	"github.com/JResearchLabs/Flutechain/command/genesis"
	"github.com/JResearchLabs/Flutechain/command/helper"
	"github.com/JResearchLabs/Flutechain/command/ibft"
	"github.com/JResearchLabs/Flutechain/command/license"
	"github.com/JResearchLabs/Flutechain/command/loadbot"
	"github.com/JResearchLabs/Flutechain/command/monitor"
	"github.com/JResearchLabs/Flutechain/command/peers"
	"github.com/JResearchLabs/Flutechain/command/secrets"
	"github.com/JResearchLabs/Flutechain/command/server"
	"github.com/JResearchLabs/Flutechain/command/status"
	"github.com/JResearchLabs/Flutechain/command/txpool"
	"github.com/JResearchLabs/Flutechain/command/version"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Flutechain is a framework for building Ethereum-compatible Blockchain networks",
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
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
