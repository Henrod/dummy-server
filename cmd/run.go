// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/Henrod/dummy-server/app"
	"github.com/Henrod/dummy-server/parser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "start dummy server",
	Long:  `starts dummy server by reading config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server")
		readConfig()

		log.Print("getting parser")
		parser := parser.NewConfigParser()

		log.Print("reading config")
		app, err := app.NewApp(parser)
		if err != nil {
			log.Fatal(err)
		}

		log.Print("starting app")
		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
