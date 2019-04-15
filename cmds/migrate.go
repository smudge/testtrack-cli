package cmds

import (
	"github.com/Betterment/testtrack-cli/migrations"
	"github.com/spf13/cobra"
)

var migrateDoc = `
Runs all migrations that haven't been applied yet in this ecosystem.
`

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run outstanding migrations",
	Long:  migrateDoc,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return migrate()
	},
}

func migrate() error {
	runner, err := migrations.NewRunner()
	if err != nil {
		return err
	}

	err = runner.RunOutstanding()
	if err != nil {
		return err
	}

	return nil
}
