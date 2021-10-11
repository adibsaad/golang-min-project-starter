package cmd

import (
	"fmt"
	"my-app/internal/settings"
	"time"

	"github.com/spf13/cobra"
)

func getProcessCommand(config *settings.Config) *cobra.Command {
	command := cobra.Command{
		Use:               "process",
		Short:             "Do some kind of processing",
		PersistentPreRunE: genPreRunE(config)([]configFn{initDatabase}),
		RunE: func(cmd *cobra.Command, args []string) error {
			// ...
			// Do some proecessing
			// ...
			fmt.Println("Processing...")
			time.Sleep(time.Second)
			fmt.Println("Processed!")

			return nil
		},
	}

	return &command
}
