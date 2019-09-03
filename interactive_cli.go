package parking

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
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

	var line string
	var firstline = true
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return err
		}
		command := parseCommand(line)
		if firstline {
			firstline = false
			if Input(command[0]) != CreateParking {
				return errors.New("First command has to be create parking")
			}
			maxSlots, err := strconv.Atoi(command[1])
			if err != nil {
				return err
			}
			p := NewParking(maxSlots)
			resp, err := processCommand(p, command)
			if err != nil {
				return err
			}
			log.Println(resp)
		}
	}
	return nil
}

// ExecuteCLIruns a CLI to execute commands.
func ExecuteCLI() {

}
