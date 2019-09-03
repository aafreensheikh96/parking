package parking

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ExecuteFile executes the input file commands.
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
	p := NewParking(0)
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimRight(line, "\n")
		command := parseCommand(line)
		if firstline {
			if Input(command[0]) != CreateParking {
				return errors.New("First command has to be create parking")
			}
			maxSlots, err := strconv.Atoi(command[1])
			if err != nil {
				return err
			}
			p = NewParking(maxSlots)
			resp, err := processCommand(p, command, firstline)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(resp)
			}
			firstline = false
			continue
		}
		resp, err := processCommand(p, command, firstline)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	}
}

// InteractiveSession a CLI to execute commands.
func InteractiveSession() error {
	var (
		command  = ""
		commands []string
		p        = NewParking(0)
		first    = true
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nInput")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	commands = parseCommand(text)
	if Input(commands[0]) != CreateParking {
		fmt.Println("\nOutput\nfirst command needs to be creating the storey")
	} else {
		maxSlots, err := strconv.Atoi(commands[1])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			p = NewParking(maxSlots)
			fmt.Println("\nOutput")
			resp, err := processCommand(p, commands, true)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				first = false
				fmt.Println(resp)
			}
		}
	}

	for command != "Exit" {
		fmt.Println("\nInput")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		commands = parseCommand(text)
		resp, err := processCommand(p, commands, first)
		fmt.Println("\nOutput")
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			fmt.Println(resp)
		}
		command = commands[0]
	}
	return nil
}
