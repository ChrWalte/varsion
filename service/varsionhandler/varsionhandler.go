package varsionhandler

import (
	"strconv"
	"strings"
	"varsion/data/fileio"
	"varsion/domain/objects/varsion"
	"varsion/service/configuration"
)

// the read varsion file data.
var loadedVarsion *varsion.Varsion

// the path to the loaded varsion file.
var loadedVarsionPath string

// get the varsion as string.
func GetVarsionString() string {
	// return loaded varsion as string.
	return loadedVarsion.ToString()
}

// Read the varsion file at the given location using the given configuration.
// pathToFile: the path to the varsion file to read.
// config: the configuration to use when handling varsions.
func ReadVarsionFile(pathToFile string, config configuration.VarsionConfig) error {
	// read the varsion data as a string.
	varsionString, err := fileio.ReadFileAsString(pathToFile)

	// if an error happened, return.
	if err != nil {
		return err
	}

	// split the varsion data on the config delimiter.
	varsionSplit := strings.Split(varsionString, config.Delimiter)

	// major will be the first split with the prefix chopped off.
	majorString := strings.TrimPrefix(varsionSplit[0], config.Prefix)

	// minor will be the second split.
	minorString := varsionSplit[1]

	// patch will be the third split with the suffix chopped off.
	patchString := strings.TrimSuffix(varsionSplit[2], config.Suffix)

	// try-parse as uint64
	major, err := strconv.ParseUint(majorString, 10, 64)
	// if an error happened, return.
	if err != nil {
		return err
	}

	// try-parse as uint64
	minor, err := strconv.ParseUint(minorString, 10, 64)
	// if an error happened, return.
	if err != nil {
		return err
	}

	// try-parse as uint64
	patch, err := strconv.ParseUint(patchString, 10, 64)
	// if an error happened, return.
	if err != nil {
		return err
	}

	// set the loaded varsion file path.
	loadedVarsionPath = pathToFile

	// set the loaded varsion.
	loadedVarsion = &varsion.Varsion{
		Prefix:    config.Prefix,
		Major:     major,
		Minor:     minor,
		Patch:     patch,
		Suffix:    config.Suffix,
		Delimiter: config.Delimiter,
	}

	// return no error.
	return nil
}

// Create a new varsion and save it to the given path.
// pathToFile: the path where the new varsion file should be created.
// config: the configuration to use with the varsion data.
func CreateVarsionFile(pathToFile string, config configuration.VarsionConfig) error {
	// create a new varsion.
	varsion := varsion.InitializeVarsion(config.Prefix, config.Suffix, config.Delimiter)

	// set the varsion data.
	loadedVarsion = &varsion

	// set the varsion path.
	loadedVarsionPath = pathToFile

	// return result of rewriting the vatsion file.
	return rewriteVarsionFile()
}

// increment the major segment
func IncrementMajor() error {
	// increment the major segment
	varsion := loadedVarsion.IncrementMajor()

	// store the new varsion.
	loadedVarsion = &varsion

	// return result of rewriting the vatsion file.
	return rewriteVarsionFile()
}

// increment the minor segment
func IncrementMinor() error {
	// increment the minor segment
	varsion := loadedVarsion.IncrementMinor()
	// store the new varsion.
	loadedVarsion = &varsion
	// return result of rewriting the vatsion file.
	return rewriteVarsionFile()
}

// increment the patch segment
func IncrementPatch() error {
	// increment the patch segment
	varsion := loadedVarsion.IncrementPatch()
	// store the new varsion.
	loadedVarsion = &varsion
	// return result of rewriting the vatsion file.
	return rewriteVarsionFile()
}

// attempt to rewrite the varsion data to the saved file location.
func rewriteVarsionFile() error {
	// return result of rewriting the vatsion file.
	return fileio.WriteFileWithString(loadedVarsionPath, loadedVarsion.ToString())
}
