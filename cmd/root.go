/*
Copyright © 2022 uscan team

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/Ankr-network/uscan/pkg"
	"os"

	"github.com/Ankr-network/uscan/share"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uscan",
	Short: "boot blockchain scan",
	Long:  ``,
	Run:   pkg.MainRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .uscan.yaml)")

	rootCmd.Flags().StringP(share.HttpAddr, "", "0.0.0.0", "service boot with this address")
	rootCmd.Flags().StringP(share.HttpPort, "", "4322", "service boot with this address")
	rootCmd.Flags().StringSliceP(share.RpcUrls, "", []string{}, "get data from blockchain, use wsurl")
	rootCmd.Flags().Uint64P(share.WorkChan, "", 24, "Open multiple works to get data")
	rootCmd.Flags().StringP(share.MdbxPath, "", "uscandb", "mdbx path")
	rootCmd.Flags().Uint64P(share.ForkBlockNum, "", 12, "fork block number")

	rootCmd.Flags().StringP(share.APPTitle, "", "", "app_title is a user-defined browser title, such as Coq, which displays Coq Chain Scan")
	rootCmd.Flags().StringP(share.UnitDisplay, "", "", "unit_display indicates the unit specified by the user, for example, Eth, Peel, Bnb")
	rootCmd.Flags().StringP(share.NodeUrl, "", "", "The node_url is the node to be used when you need to interact with the contract")
	rootCmd.Flags().Uint64P(share.Decimal, "", 18, "decimal")

	// bind viper
	viper.BindPFlag(share.HttpAddr, rootCmd.Flags().Lookup(share.HttpAddr))
	viper.BindPFlag(share.HttpPort, rootCmd.Flags().Lookup(share.HttpPort))
	viper.BindPFlag(share.RpcUrls, rootCmd.Flags().Lookup(share.RpcUrls))
	viper.BindPFlag(share.WorkChan, rootCmd.Flags().Lookup(share.WorkChan))
	viper.BindPFlag(share.MdbxPath, rootCmd.Flags().Lookup(share.MdbxPath))
	viper.BindPFlag(share.ForkBlockNum, rootCmd.Flags().Lookup(share.ForkBlockNum))

	viper.BindPFlag(share.APPTitle, rootCmd.Flags().Lookup(share.APPTitle))
	viper.BindPFlag(share.UnitDisplay, rootCmd.Flags().Lookup(share.UnitDisplay))
	viper.BindPFlag(share.NodeUrl, rootCmd.Flags().Lookup(share.NodeUrl))
	viper.BindPFlag(share.Decimal, rootCmd.Flags().Lookup(share.Decimal))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".uscan" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".uscan")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
