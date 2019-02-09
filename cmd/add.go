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
	"fmt"
	"strconv"
	"time"

	"github.com/edznux/metastats/config"
	"github.com/edznux/metastats/events"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A new value(s) to the dataset passed in the first arguments.",
	Example: `metastats add steps 1337
	will add 1337 and the current timestamp in the file steps.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		fileName := cfg.DataPath + args[0] + ".csv"

		date := time.Now().Unix()

		data := []string{
			strconv.FormatInt(date, 10),
		}
		data = append(data, args[1:]...)

		if cfg.Verbose {
			fmt.Println("Saving data : [", data, "] to file : ", fileName)
		}

		err := events.SaveToFile(fileName, data)
		if err != nil {
			fmt.Printf("Could not save this new event : %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
