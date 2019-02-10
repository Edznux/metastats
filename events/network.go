package events

import (
	"bufio"
	"fmt"
	"github.com/edznux/metastats/config"
	"log"
	"os"
	"strconv"
	"time"
)

func readNetworkFromSys(config config.Config) []string {
	stats := []string{"rx_bytes", "tx_bytes"}

	res := []string{}
	for _, stat := range stats {
		path := fmt.Sprintf("/sys/class/net/%s/statistics/%s", config.Interface, stat)

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// read the first line. No need loop
		scanner.Scan()
		data := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalf("Cannot read %s : %s", path, err)
		}
		res = append(res, data)
	}

	return res
}

func formatNetwork(config config.Config) []string {
	data := readNetworkFromSys(config)
	date := time.Now().Unix()

	res := []string{
		strconv.FormatInt(date, 10),
	}
	res = append(res, data...)

	return res
}
