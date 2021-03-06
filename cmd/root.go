// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "credscontroller",
	Short: "creds controller manages credentials for kubernetes applications",
	Long:  "creds controller helps building  a secure workflow for suppling credentials to containers running in openshift, currently it integrates with Vault",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	//fmt.Println("root.init")
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().String("vault-addr", "https://vault:8200", "Vault URL")
	RootCmd.PersistentFlags().String("vault-cacert", "/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt", "ca certificate to be used to validate the connection to Vault")
	RootCmd.PersistentFlags().String("log-level", "info", "log level")
	viper.BindPFlag("vault-addr", RootCmd.PersistentFlags().Lookup("vault-addr"))
	viper.BindPFlag("vault-cacert", RootCmd.PersistentFlags().Lookup("vault-cacert"))
	viper.BindPFlag("log-level", RootCmd.PersistentFlags().Lookup("log-level"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match
}
