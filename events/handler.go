package events

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
)

var wg sync.WaitGroup

const (
	verbose            = true
	micePath           = "/dev/input/mouse0"
	keyboardPath       = "/dev/input/event4"
	saveTimer          = 5 * time.Second
	SavingPath         = "./data/"
	savingPathMice     = SavingPath + "mice.dat"
	savingPathKeyboard = SavingPath + "keyboard.dat"
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

func Start() {

	var data *InputData
	data = &InputData{}

	ticker := time.NewTicker(saveTimer)

	wg.Add(1)
	go saveAndReset(data, ticker)
	go monitorMice(data, ticker)
	go monitorKeyboard(data, ticker)
	wg.Wait()
}

func monitorMice(inputs *InputData, ticker *time.Ticker) {
	device, _ := evdev.Open(micePath)
	for {
		ie, err := device.Read()
		if err != nil {
			fmt.Println("Error :", err)
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

	device, _ := evdev.Open(keyboardPath)
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
	for {
		select {
		case <-ticker.C:
			if verbose {
				fmt.Printf("click count | Total: %d, Left : %d, Middle : %d, Right : %d\n", (data.Mice.LeftCount + data.Mice.MiddleCount + data.Mice.RightCount), data.Mice.LeftCount, data.Mice.MiddleCount, data.Mice.RightCount)
				fmt.Printf("keyboard press | Total: %d\n", data.Keyboard.Keypress)
			}

			dataMice := formatMiceOutput(data)
			dataKb := formatKbOutput(data)
			SaveToFile(savingPathMice, dataMice)
			SaveToFile(savingPathKeyboard, dataKb)
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

func SaveToFile(filename string, data []string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()
}
