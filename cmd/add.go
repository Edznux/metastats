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
