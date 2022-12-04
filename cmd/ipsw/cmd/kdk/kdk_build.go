package kdk

import (
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	KDKCmd.AddCommand(kdkBuildCmd)

	kdkBuildCmd.Flags().StringP("kernelcache", "c", "", "path to the decompressed kernelcache to source kexts from")
	kdkBuildCmd.Flags().StringP("kernel", "k", "", "path to the kernel to source symbolsets from")
	kdkBuildCmd.Flags().String("kdk", "o", "path to place assembled the KDK")

	kdkBuildCmd.MarkFlagRequired("kernel")
	kdkBuildCmd.MarkFlagRequired("kernelcache")
	kdkBuildCmd.MarkFlagRequired("kdk")

	viper.BindPFlag("kdk.build.kernel", kdkBuildCmd.Flags().Lookup("kernel"))
	viper.BindPFlag("kdk.build.kernelcache", kdkBuildCmd.Flags().Lookup("kernelcache"))
	viper.BindPFlag("kdk.build.kdk", kdkBuildCmd.Flags().Lookup("kdk"))
}

var kdkBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a KDK",
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}
		// builder, err := NewKDKBuilderFromPath(args[0])
		builder := &KDKBuilder{
			KernelcachePath:    viper.GetString("kdk.build.kernelcache"),
			KernelPath:         viper.GetString("kdk.build.kernel"),
			DestinationKDKPath: viper.GetString("kdk.build.kdk"),
		}
		return builder.BuildKDK()
	},
}
