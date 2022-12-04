package kdk

import (
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	KDKCmd.AddCommand(kdkKernelcacheCmd)

	kdkKernelcacheCmd.MarkZshCompPositionalArgumentFile(1, "*.yaml")

	kdkKernelcacheCmd.Flags().StringP("kernel", "k", "", "the path of the kernel to embed in the kernelcache")
	kdkKernelcacheCmd.Flags().StringP("output", "o", "", "the path to write the kernelcache to")
	kdkKernelcacheCmd.Flags().String("kdk", "", "the path to the KDK to pass to kmutil")

	kdkKernelcacheCmd.MarkFlagRequired("kernel")
	kdkKernelcacheCmd.MarkFlagRequired("output")
	kdkKernelcacheCmd.MarkFlagRequired("kdk")

	viper.BindPFlag("kdk.build-kc.kernel", kdkKernelcacheCmd.Flags().Lookup("kernel"))
	viper.BindPFlag("kdk.build-kc.output", kdkKernelcacheCmd.Flags().Lookup("output"))
	viper.BindPFlag("kdk.build-kc.kdk", kdkKernelcacheCmd.Flags().Lookup("kdk"))
}

var kdkKernelcacheCmd = &cobra.Command{
	Use:   "build-kc",
	Short: "Builds a kernelcache using a custom KDK built by ipsw",
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		builder := &KDKBuilder{
			KernelPath:                 viper.GetString("kdk.build-kc.kernel"),
			DestinationKernelcachePath: viper.GetString("kdk.build-kc.output"),
			DestinationKDKPath:         viper.GetString("kdk.build-kc.kdk"),
		}

		return builder.CreateKernelcache()
	},
}
