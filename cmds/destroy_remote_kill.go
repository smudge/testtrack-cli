package cmds

import (
	"github.com/Betterment/testtrack-cli/migrationmanagers"
	"github.com/Betterment/testtrack-cli/remotekills"
	"github.com/spf13/cobra"
)

var destroyRemoteKillDoc = `
Unsets a remote kill, allowing users of affected apps to see whatever variant
of the split they would otherwise see.

Example:

testtrack unset_remote_kill my_fancy_experiment
`

func init() {
	destroyCmd.AddCommand(destroyRemoteKillCmd)
}

var destroyRemoteKillCmd = &cobra.Command{
	Use:   "remote_kill split_name reason override_to first_bad_version [fixed_version]",
	Short: "Remove a split remote-kill for a range of app versions",
	Long:  destroyRemoteKillDoc,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return destroyRemoteKill(&args[0], &args[1])
	},
}

func destroyRemoteKill(split, reason *string) error {
	remoteKill, err := remotekills.New(split, reason, nil, nil, nil)
	if err != nil {
		return err
	}

	mgr, err := migrationmanagers.New(remoteKill)
	if err != nil {
		return err
	}

	err = mgr.Save()
	if err != nil {
		return err
	}

	return nil
}