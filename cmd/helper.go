package cmd

import (
	"github.com/spf13/cobra"
	"my-app/internal/settings"
)

type configFn func(*settings.Config) error

// Returns a function that executes the all
// th passed-in functions in sequential order
// (passing the config object to each one)
// before the main command is run.
func genPreRunE(config *settings.Config) func([]configFn) func(*cobra.Command, []string) error {
	return func(fns []configFn) func(*cobra.Command, []string) error {
		return func(command *cobra.Command, strings []string) error {
			for i := 0; i < len(fns); i++ {
				fn := fns[i]
				if err := fn(config); err != nil {
					return err
				}
			}

			return nil
		}
	}
}

func initDatabase(config *settings.Config) error {
	// ...
	// Initiate a connection to the db to make sure it's available.
	// ...

	return nil
}
