package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"varsion/data/fileio"
	"varsion/service/configuration"
	"varsion/service/varsionhandler"
)

var givenCommands []string = make([]string, 0)

func main() {
	workingPath := getCurrentWorkingPath()
	pathToConfig := path.Join(workingPath, ".atoolconfig.json")
	isConfigLoaded := getConfiguration(pathToConfig)

	if !isConfigLoaded {
		fmt.Println("[INFO]: configuration not loaded")
		os.Exit(1)
	}

	varsionConfig := configuration.Data.Varsion
	pathToVarsion := path.Join(workingPath, varsionConfig.FileName)
	isVarsionLoaded := getVarsionFile(pathToVarsion, varsionConfig)

	if !isVarsionLoaded {
		fmt.Println("[WARN]: version not loaded")
		os.Exit(1)
	}

	isCommandsLoaded := getCommandsFromArgs()

	if !isCommandsLoaded {
		fmt.Println("[WARN]: commands not given")
		os.Exit(1)
	}

	handleCommands(pathToVarsion)
}

// returns where the binary file was called from
// this is the directory where the VERSION and .atoolconfig.json files should be
// if this fails, the application fails
func getCurrentWorkingPath() string {
	workingPath, err := os.Getwd()
	if err != nil {
		fmt.Println("[ERROR]: failed getting working directory")
		log.Fatal(err)
	}
	return workingPath
}

func getConfiguration(pathToConfig string) bool {
	err := configuration.Initialize(pathToConfig)
	if err != nil {
		fmt.Println("[WARN]: configuration failed to load, attempting to create")
		return createConfiguration(pathToConfig)
	}
	return true
}

func createConfiguration(pathToConfig string) bool {
	newConfiguration := configuration.Config{
		Varsion: configuration.VarsionConfig{
			FileName:  "VERSION",
			Prefix:    "",
			Suffix:    "",
			Delimiter: ".",
		},
	}

	configurationJson, err := json.MarshalIndent(newConfiguration, "", "\t")
	if err != nil {
		fmt.Println("[ERROR]: configuration failed to create")
		log.Fatal(err)
	}

	err = fileio.WriteFile(pathToConfig, configurationJson)
	if err != nil {
		fmt.Println("[ERROR]: configuration failed to create")
		log.Fatal(err)
	}

	err = configuration.Initialize(pathToConfig)
	if err != nil {
		fmt.Println("[ERROR]: configuration failed to create")
		log.Fatal(err)
	}

	return true
}

func getVarsionFile(pathToFile string, config configuration.VarsionConfig) bool {
	err := varsionhandler.ReadVarsionFile(pathToFile, config)
	if err != nil {
		fmt.Println("[INFO]: version failed to load, attempting to create")
		fmt.Println(err)
		return createVarsionFile(pathToFile, config)
	}
	return true
}

func createVarsionFile(pathToFile string, config configuration.VarsionConfig) bool {
	err := varsionhandler.CreateVarsionFile(pathToFile, config)
	if err != nil {
		fmt.Println("[ERROR]: version failed to create")
		log.Fatal(err)
	}
	return true
}

func getCommandsFromArgs() bool {
	if len(os.Args) >= 2 {
		givenCommands = os.Args[1:]
		return true
	}
	return getCommandsFromUser()
}

func getCommandsFromUser() bool {
	var input string

	currentVarsion := varsionhandler.GetVarsionString()
	fmt.Printf("[INFO]: current version: [%s]\n", currentVarsion)
	fmt.Print("[INPUT]: segment to change (major, minor, patch): ")
	fmt.Scanln(&input)
	givenCommands = append(givenCommands, input)

	// fmt.Print("increment or decrement (+, -): ")
	// fmt.Scanln(input)
	// givenCommands = append(givenCommands, input)
	return true
}

func handleCommands(pathToFile string) {
	var err error

	varsionFrom := varsionhandler.GetVarsionString()

	givenCommand := strings.ToUpper(givenCommands[0])

	switch givenCommand {
	case "MAJOR":
		err = varsionhandler.IncrementMajor()
	case "MA":
		err = varsionhandler.IncrementMajor()
	case "MINOR":
		err = varsionhandler.IncrementMinor()
	case "MI":
		err = varsionhandler.IncrementMinor()
	case "PATCH":
		err = varsionhandler.IncrementPatch()
	case "P":
		err = varsionhandler.IncrementPatch()
	default:
		log.Fatal("[ERROR]: unknown command")
	}

	if err != nil {
		fmt.Println("[ERROR]: version failed to update")
		log.Fatal(err)
	}

	varsionTo := varsionhandler.GetVarsionString()
	fmt.Printf("[INFO]: [%s] -> [%s]", varsionFrom, varsionTo)
}
