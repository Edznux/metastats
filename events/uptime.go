package events

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readUptimeFromProc() string {

	file, err := os.Open("/proc/uptime")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read the first line. No need loop
	scanner.Scan()
	data := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal("Cannot read /proc/uptime file", err)
	}

	return data
}

func formatUptime() []string {
	data := readUptimeFromProc()
	splitted := strings.Split(data, " ")

	fmt.Printf("splitted : %s", splitted)
	return splitted
}
