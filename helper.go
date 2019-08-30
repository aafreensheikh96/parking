package parking

import (
	"strings"
)

func parseCommand(command string) []string {
	parsedCommand := []string{}

	command = strings.Replace(command, Tab, Space, -1)

	for _, s := range strings.Split(command, Space) {
		if s != "" {
			parsedCommand = append(parsedCommand, s)
		}
	}

	return parsedCommand
}

// processCommand process each command
func processCommand(pl *ParkingLotService, command []string) (*Resp, error) {
	return nil, nil
}
