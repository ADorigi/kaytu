package plugin

import (
	"errors"
	"github.com/kaytu-io/kaytu/pkg/plugin"
	"github.com/spf13/cobra"
	"os"
)

var uninstallCmd = &cobra.Command{
	Use: "uninstall",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := plugin.New()
		err := manager.StartServer()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			return errors.New("please provide plugin name")
		}

		err = manager.Uninstall(args[0])
		if err != nil {
			os.Stderr.WriteString("failed to uninstall plugin due to %v\n", err)
			return err
		}
		return nil
	},
}
