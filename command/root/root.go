package root

import (
	"fmt"
	"os"

	"github.com/JResearchLabs/aigisos/command/backup"
	"github.com/JResearchLabs/aigisos/command/genesis"
	"github.com/JResearchLabs/aigisos/command/helper"
	"github.com/JResearchLabs/aigisos/command/ibft"
	"github.com/JResearchLabs/aigisos/command/license"
	"github.com/JResearchLabs/aigisos/command/loadbot"
	"github.com/JResearchLabs/aigisos/command/monitor"
	"github.com/JResearchLabs/aigisos/command/peers"
	"github.com/JResearchLabs/aigisos/command/secrets"
	"github.com/JResearchLabs/aigisos/command/server"
	"github.com/JResearchLabs/aigisos/command/status"
	"github.com/JResearchLabs/aigisos/command/txpool"
	"github.com/JResearchLabs/aigisos/command/version"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "aigisos is a framework for building Ethereum-compatible Blockchain networks",
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
