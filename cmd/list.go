package cmd

import (
	"fmt"
	"github.com/edznux/metastats/config"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cfg config.Config

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the possible events",
	Run: func(cmd *cobra.Command, args []string) {
		cfg = config.LoadConfig()
		events := ListEventsName()
		fmt.Println("Available events :")
		for _, ev := range events {
			fmt.Println(" -", ev)
		}
	},
}

func ListEventsName() []string {
	eventNameList := []string{}
	var fullPath string
	var eventName string
	files, err := ioutil.ReadDir(cfg.DataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fullPath = path.Join(cfg.DataPath, file.Name())
		if filepath.Ext(fullPath) == ".csv" {
			eventName = strings.Replace(file.Name(), ".csv", "", -1)
		}
		switch eventName {
		// multi line case
		case "keyboard", "mice", "network":
			eventName = "[default] " + eventName
			break
		// non default events
		default:
			eventName = "[custom] " + eventName
			break
		}

		eventNameList = append(eventNameList, eventName)
	}
	return eventNameList
}

func init() {
	rootCmd.AddCommand(listCmd)
}
