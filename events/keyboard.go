package events

import (
	"fmt"
	"strconv"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
)

func monitorKeyboard(inputs *InputData) {

	device, _ := evdev.Open(Cfg.KeyboardPath)
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

func formatKbOutput(data *InputData) []string {

	date := time.Now().Unix()
	return []string{
		strconv.FormatInt(date, 10),
		strconv.FormatInt(int64(data.Keyboard.Keypress), 10),
	}
}
