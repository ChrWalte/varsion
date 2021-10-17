package fileio

import (
	"os"
)

// the file permissions for read, write, and execute.
const filePermissions = 0777

// Read the file at the given path and return the contents as bytes.
// pathToFile: path to the file to read.
func ReadFile(pathToFile string) ([]byte, error) {
	// get file contents
	contents, err := os.ReadFile(pathToFile)

	// return contents and err
	return contents, err
}

// Read the file at the given path and return the contents as a string.
// pathToFile: path to the file to read.
func ReadFileAsString(pathToFile string) (string, error) {
	// read the file
	contentBytes, err := ReadFile(pathToFile)

	// convert to string.
	content := string(contentBytes)

	// return contents and err
	return content, err
}

// Write the contents to the file at the given path.
// pathToFile: path to the file to write.
// contents: data to write to the file.
func WriteFile(pathToFile string, contents []byte) error {
	// write to the file.
	err := os.WriteFile(pathToFile, contents, filePermissions)

	// return err.
	return err
}

// Write the contents as a string to the file at the given path.
// pathToFile: path to the file to write.
// contents: data to write to the file.
func WriteFileWithString(filePath string, contents string) error {
	// convet from string to bytes.
	contentBytes := []byte(contents)

	// write to file using bytes.
	return WriteFile(filePath, contentBytes)
}
