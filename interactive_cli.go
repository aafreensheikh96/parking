package parking

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ExecuteFile, executes the input file commands.
func ExecuteFile(filename string) error {
	file, err := os.Open(strings.Trim(filename, " "))
	defer file.Close()
	if err != nil {
		log.Print(err)
		return err
	}

	reader := bufio.NewReader(file)

	// var line string
	for {
		_, err = reader.ReadString('\n')
		if err != nil {
			return err
		}
	}
	return nil
}

// ExecuteCLIruns a CLI to execute commands.
func ExecuteCLI() {}
