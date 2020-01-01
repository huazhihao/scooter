// Copyright © 2019 Hua Zhihao <ihuazhihao@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/huazhihao/scooter/pkg/commons"
	"github.com/huazhihao/scooter/pkg/log"
)

var (
	cfgFile   string
	debugMode bool
	config    commons.Config
)

var rootCmd = &cobra.Command{
	Use:   "scooter",
	Short: "A fast and fully featured reverse proxy",
	Long: `🛵 Scooter is a lightweight L4+L7 reverse proxy and load balancer written in Go. It provides nginx-like functionalities with little effort on setup, and better integration with modern monitoring tools..

  Find more information at: https://github.com/huazhihao/scooter`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if debugMode {
			log.SetLevel("debug")
		}

		cfgFile = viper.ConfigFileUsed()
		log.Debugf("Using config file %s:", cfgFile)
		// if err := viper.ReadInConfig(); err != nil {
		// 	log.Fatalf("Error while reading config file %s: %v", cfgFile, err)
		// }
		// err := viper.Unmarshal(&config)

		data, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.Fatalf("Error while reading %s: %v", cfgFile, err)
		}

		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatalf("unable to decode config file: %v", err)
		}
		log.Debugf("%+v", config)
	},

	Run: func(cmd *cobra.Command, args []string) {
		s := commons.NewScooter(config)
		s.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./scooter.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "turn on debug log")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Errorf("Error while finding home directory: %v", err)
			os.Exit(1)
		}
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/scooter/")
		viper.SetConfigName("scooter")
	}

	// viper.AutomaticEnv()

}
