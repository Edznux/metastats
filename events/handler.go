package events

import (
	"encoding/csv"
	"fmt"
	"github.com/edznux/metastats/config"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	//Cfg is the global config variable. May be overriden by cmd/*
	Cfg config.Config
)

type InputData struct {
	Timestamp time.Time
	Mice      Mice
	Keyboard  Keyboard
	Uptime    []string
}

type Keyboard struct {
	Keypress int
}

type Mice struct {
	LeftCount   int
	RightCount  int
	MiddleCount int
}

func init() {
	loadConfig()
}
func loadConfig() {
	Cfg = config.LoadConfig()
	err := os.MkdirAll(Cfg.DataPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create data path %s", err)
	}
	err = os.MkdirAll(Cfg.LogPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create log path %s", err)
	}
}

func Start() {
	var data *InputData
	data = &InputData{}

	ticker := time.NewTicker(time.Duration(Cfg.SaveTimer) * time.Second)

	wg.Add(1)
	go saveAndReset(data, ticker)
	go monitorMice(data)
	go monitorKeyboard(data)

	wg.Wait()
}

func saveAndReset(data *InputData, ticker *time.Ticker) {
	var err error

	for {
		select {
		case <-ticker.C:
			if Cfg.Verbose {
				log.Printf("click count | Total: %d, Left : %d, Middle : %d, Right : %d\n", (data.Mice.LeftCount + data.Mice.MiddleCount + data.Mice.RightCount), data.Mice.LeftCount, data.Mice.MiddleCount, data.Mice.RightCount)
				log.Printf("keyboard press | Total: %d\n", data.Keyboard.Keypress)
			}

			dataMice := formatMiceOutput(data)
			dataKb := formatKbOutput(data)
			uptime := formatUptime()
			network := formatNetwork(Cfg)

			err = SaveToFile(filepath.Join(Cfg.DataPath, "mice.csv"), dataMice)
			if err != nil {
				log.Printf("Could not save to mice file : %s\n", err)
			}
			err = SaveToFile(filepath.Join(Cfg.DataPath, "keyboard.csv"), dataKb)
			if err != nil {
				log.Printf("Could not save to keyboard file : %s\n", err)
			}
			err = SaveToFile(filepath.Join(Cfg.DataPath, "uptime.csv"), uptime)
			if err != nil {
				log.Printf("Could not save to uptime file : %s\n", err)
			}
			err = SaveToFile(filepath.Join(Cfg.DataPath, "network.csv"), network)
			if err != nil {
				log.Printf("Could not save to network file : %s\n", err)
			}

			// reset count !
			data.Mice = Mice{}
			data.Keyboard = Keyboard{}
		}
	}
}

func SaveToFile(filename string, data []string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("Could not save to file : %s", err)
	}
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()

	return nil
}
