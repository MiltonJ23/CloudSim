/*
Copyright © 2025 Zingui Fred Mike mikezingui@yahooc.com
*/
package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudsim",
	Short: "A cloud software as a service simulation component",
	Long: `cloudsim is a CLI tool for simulating a Cloud SaaS environment.

This application allows you to model and interact with virtual cloud infrastructure
components. The primary component is the StorageVirtualNode, which simulates
the behavior of a cloud storage service.

Key simulation features include:
* Tracking storage capacity and usage.
* Managing network bandwidth and utilization.
* Simulating realistic, chunk-based file transfer  operations.
* Providing metrics for resource utilization.

This tool can be used to test storage management strategies, network performance,
or application behavior in a controlled, simulated cloud environment.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithNotifySignal(os.Interrupt, os.Kill),
	); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CloudSim.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
