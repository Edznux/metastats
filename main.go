package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

const micePath = "/dev/input/mouse0"
const saveTimer = 5 * time.Second
const savingPath = "./data/events.dat"

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

func main() {
	var data *InputData
	data = &InputData{}

	ticker := time.NewTicker(saveTimer)

	wg.Add(1)
	go saveAndReset(data, ticker)
	go monitorMice(data, ticker)
	wg.Wait()
}

func monitorMice(inputs *InputData, ticker *time.Ticker) {

	file, _ := os.Open(micePath)

	for {
		data := make([]byte, 3)
		byteRead, err := file.Read(data)

		if err != nil {
			log.Fatal(err)
		}

		if byteRead < 3 {
			fmt.Println("Error : byteRead < 3 ! (", byteRead, ")")
			continue
		}

		switch data[0] {
		case 0x9:
			// fmt.Println("LEFT Click !")
			inputs.Mice.LeftCount++
		case 0xa:
			// fmt.Println("RIGHT Click !")
			inputs.Mice.RightCount++
		case 0xc:
			// fmt.Println("MIDDLE Click !")
			inputs.Mice.MiddleCount++
		}
	}
}

func saveAndReset(data *InputData, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			fmt.Printf("click count | Total: %d, Left : %d, Middle : %d, Right : %d\n", (data.Mice.LeftCount + data.Mice.MiddleCount + data.Mice.RightCount), data.Mice.LeftCount, data.Mice.MiddleCount, data.Mice.RightCount)

			f, err := os.OpenFile(savingPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

			if err != nil {
				fmt.Println("Error : ", err.Error())
			}
			defer f.Close()

			w := bufio.NewWriter(f)

			_, err = fmt.Fprintf(w, "%d,%d,%d,%d\n", (data.Mice.LeftCount + data.Mice.MiddleCount + data.Mice.RightCount), data.Mice.LeftCount, data.Mice.MiddleCount, data.Mice.RightCount)

			if err != nil {
				fmt.Println("Error : ", err.Error())
			}
			err = w.Flush()

			if err != nil {
				fmt.Println("Error : ", err.Error())
			}

			// reset count !
			data.Mice = Mice{}
			data.Keyboard = Keyboard{}
		}
	}
}
