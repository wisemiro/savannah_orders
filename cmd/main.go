package main

import (
	"os"
	"savannah/cmd/app"

	"github.com/spf13/cobra"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func Serve() *cobra.Command {
	serverCmd := &cobra.Command{
		Use: "start",
	}
	serverCmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		server, err := app.StartApp(cfg)
		if err != nil {
			return err
		}

		return server.ServerConnection()
	}
	return serverCmd
}

func Run(args []string) error {
	rootCmd := &cobra.Command{
		Use: "Savanah orders",
	}
	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file to use.")
	rootCmd.AddCommand(Serve())
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}
