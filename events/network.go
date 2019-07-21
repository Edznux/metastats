package events

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/edznux/metastats/config"
	"github.com/pkg/errors"
)

func readNetworkFromSys(config config.Config) ([]string, error) {
	stats := []string{"rx_bytes", "tx_bytes"}

	res := []string{}
	for _, stat := range stats {
		path := fmt.Sprintf("/sys/class/net/%s/statistics/%s", config.Interface, stat)

		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		// read the first line. No need loop
		scanner.Scan()
		data := scanner.Text()
		if err := scanner.Err(); err != nil {
			errors.Wrap(err, "Cannot read "+path)
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}

func formatNetwork(config config.Config) []string {
	data, err := readNetworkFromSys(config)
	if err != nil {
		log.Print(err.Error())
	}
	date := time.Now().Unix()

	res := []string{
		strconv.FormatInt(date, 10),
	}
	res = append(res, data...)

	return res
}
