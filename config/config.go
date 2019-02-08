package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//Config defines the configuration structure for metastats
type Config struct {
	Verbose      bool   `toml:"Verbose"`
	MicePath     string `toml:"MicePath"`     // linux device path (input)
	KeyboardPath string `toml:"KeyboardPath"` // linux device path (input)
	LogPath      string `toml:"LogPath"`      // were to save the application logs (output)
	DataPath     string `toml:"DataPath"`     // were to save the collected data (output)
	SaveTimer    int    `toml:"SaveTimer"`    // save every x seconds
}

var ListOfConfigPath = []string{
	// --config /some/path will be place BEFORE any other
	"/etc/metastats/config.toml",
	// ./config.toml is added in the init function
}

func init() {
	/* add current working dir to the list of config path */
	wd, err := os.Getwd()
	if err != nil {
		log.Println("Error while getting current working directory.")
	}
	currentPath := filepath.Join(wd, "config.toml")
	ListOfConfigPath = append(ListOfConfigPath, currentPath)
}

// LoadConfig search through config files.
// It select the first existing file from the listOfConfigPath
func LoadConfig() Config {
	var config Config
	var usedPath string
	for _, path := range ListOfConfigPath {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			usedPath = path
			break
		}
	}

	if _, err := toml.DecodeFile(usedPath, &config); err != nil {
		log.Fatalln(err)
	}
	log.Println("Loaded config :", usedPath)
	return config
}
