package kdk

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var KDKCmd = &cobra.Command{
	Use:   "kdk",
	Short: "Create KDKs",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("color", cmd.Flags().Lookup("color"))
		viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
