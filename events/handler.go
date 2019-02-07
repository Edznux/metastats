package events

import (
	"encoding/csv"
	"fmt"
	"github.com/edznux/metastats/config"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
)

var (
	wg  sync.WaitGroup
	cfg config.Config
)

type InputData struct {
	Timestamp time.Time
	Mice      Mice
	Keyboard  Keyboard
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
	cfg = config.LoadConfig()
	err := os.MkdirAll(cfg.DataPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create data path %s", err)
	}
	err = os.MkdirAll(cfg.LogPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Could not create log path %s", err)
	}
}

func Start() {

	var data *InputData
	data = &InputData{}

	ticker := time.NewTicker(time.Duration(cfg.SaveTimer) * time.Second)

	wg.Add(1)
	go saveAndReset(data, ticker)
	go monitorMice(data, ticker)
	go monitorKeyboard(data, ticker)
	wg.Wait()
}

func monitorMice(inputs *InputData, ticker *time.Ticker) {
	device, _ := evdev.Open(cfg.MicePath)
	for {
		ie, err := device.Read()
		if err != nil {
			log.Println("Error :", err)
		}

		if ie[0].Time.Sec == 9 {
			inputs.Mice.LeftCount++
		}
		if ie[0].Time.Sec == 10 {
			inputs.Mice.RightCount++
		}
		if ie[0].Time.Sec == 12 {
			inputs.Mice.MiddleCount++
		}
	}
}

func monitorKeyboard(inputs *InputData, ticker *time.Ticker) {

	device, _ := evdev.Open(cfg.KeyboardPath)
	press := true
	for {
		_, err := device.Read()
		if err != nil {
			fmt.Println("Error :", err)
		}
		// handle just keypress (dismiss key release)
		if press {
			inputs.Keyboard.Keypress++
		}
		press = !press
	}
}

func saveAndReset(data *InputData, ticker *time.Ticker) {
	var err error

	for {
		select {
		case <-ticker.C:
			if cfg.Verbose {
				log.Printf("click count | Total: %d, Left : %d, Middle : %d, Right : %d\n", (data.Mice.LeftCount + data.Mice.MiddleCount + data.Mice.RightCount), data.Mice.LeftCount, data.Mice.MiddleCount, data.Mice.RightCount)
				log.Printf("keyboard press | Total: %d\n", data.Keyboard.Keypress)
			}

			dataMice := formatMiceOutput(data)
			dataKb := formatKbOutput(data)
			err = SaveToFile(filepath.Join(cfg.DataPath, "mice.dat"), dataMice)
			if err != nil {
				log.Printf("Could not save to mice file : %s\n", err)
			}
			err = SaveToFile(filepath.Join(cfg.DataPath, "keyboard.dat"), dataKb)
			if err != nil {
				log.Printf("Could not save to keyboard file : %s\n", err)
			}
			// reset count !
			data.Mice = Mice{}
			data.Keyboard = Keyboard{}
		}
	}
}
func formatMiceOutput(data *InputData) []string {

	date := time.Now().Unix()
	return []string{
		strconv.FormatInt(date, 10),
		strconv.FormatInt(int64(data.Mice.LeftCount+data.Mice.MiddleCount+data.Mice.RightCount), 10),
		strconv.FormatInt(int64(data.Mice.LeftCount), 10),
		strconv.FormatInt(int64(data.Mice.MiddleCount), 10),
		strconv.FormatInt(int64(data.Mice.RightCount), 10),
	}
}

func formatKbOutput(data *InputData) []string {

	date := time.Now().Unix()
	return []string{
		strconv.FormatInt(date, 10),
		strconv.FormatInt(int64(data.Keyboard.Keypress), 10),
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
