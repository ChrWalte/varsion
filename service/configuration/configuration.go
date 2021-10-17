package configuration

import (
	"encoding/json"
	"varsion/data/fileio"
)

// interface used for configuration
// should be updated when new configuration is added
type Config struct {

	// configuration for varsion tool
	Varsion VarsionConfig
}

// interface for the varsion configuration
// should be updated when new varsion configuration is added
type VarsionConfig struct {
	// file name of the varsion file
	FileName string

	// prefix before the varsion
	Prefix string

	// suffix after the varsion
	Suffix string

	// delimiter to split the varsion up
	Delimiter string
}

// configuration in memory
var Data Config

// reads the given configuration file and stores it into memory
// pathToConfigurationFile: path to the file to use for configuration
func Initialize(pathToConfigurationFile string) error {
	// read the config file.
	contents, err := fileio.ReadFile(pathToConfigurationFile)

	// if error return
	if err != nil {
		return err
	}

	// un-json it and store in memory
	err = json.Unmarshal(contents, &Data)

	// return err
	return err
}
