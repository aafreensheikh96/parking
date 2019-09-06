package parking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExecuteFile tests the ExecuteFile method
func TestExecuteFile(t *testing.T) {
	var inputfile = "samples/file_input.txt"

	ExecuteFile(inputfile)
}

func TestProcessCommand(t *testing.T) {
	type input struct {
		p        *parking
		commands []string
		first    bool
	}
	type output struct {
		r   *resp
		err error
	}
	tests := []struct {
		in  input
		out output
	}{
		{
			in: input{
				p:        NewParking(0),
				commands: []string{"create_parking_lot", "4"},
				first:    true,
			},
			out: output{NewResponse(NewParkingLot(4)), nil},
		},
		{
			in: input{
				p:        NewParking(0),
				commands: []string{"create_parking_lot", "4"},
				first:    false,
			},
			out: output{&resp{command: NotFisrt}, nil},
		},
	}
	assert := assert.New(t)
	for _, test := range tests {
		i := test.in
		resp, err := processCommand(i.p, i.commands, i.first)
		if err != nil {
			assert.Equal(test.out.err, err)
		}
		assert.Equal(test.out.r, resp)
	}
}
