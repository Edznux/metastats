package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/edznux/metastats/config"
	"github.com/edznux/metastats/events"
	"github.com/olebedev/when"
	"github.com/spf13/cobra"
)

var (
	at string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A new value(s) to the dataset passed in the first arguments.",
	Example: `metastats add steps 1337
	will add 1337 and the current timestamp in the file steps.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		var date int64

		cfg := config.LoadConfig()
		if len(args) == 0 {
			fmt.Println("Please provide a dataset name")
			return
		}

		fileName := cfg.DataPath + args[0] + ".csv"

		if at == "" {
			date = time.Now().Unix()
		} else {
			res, err := when.EN.Parse(at, time.Now())
			if err != nil {
				fmt.Printf("Could not parse the date : %s\n", err)
				return
			}
			date = res.Time.Unix()
		}

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
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVar(&at, "at", "", "Use a custom time. Human readable format (today 2pm for example)")
}
