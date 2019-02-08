package events

import (
	"log"
	"strconv"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
)

func monitorMice(inputs *InputData) {
	device, _ := evdev.Open(Cfg.MicePath)
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
