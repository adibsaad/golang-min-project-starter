package cmd

import (
	"my-app/internal/log"
	"my-app/internal/settings"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

func Execute() {
	var rootCmd = &cobra.Command{
		Use: "my-app",
		Run: func(cmd *cobra.Command, args []string) {
			// ...
			// Default behavior, with no command specified
			// ...
			cmd.Println(cmd.UsageString())
			os.Exit(1)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file to use")
	config := settings.NewConfig(cfgFile)

	// Add your commands here
	rootCmd.AddCommand(getProcessCommand(config))

	if err := rootCmd.Execute(); err != nil {
		log.Log.Err(err).Msg("")
		os.Exit(1)
	}
}
