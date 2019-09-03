package parking

import (
	"testing"
)

// TestExecuteFile tests the ExecuteFile method
func TestExecuteFile(t *testing.T) {
	var inputfile = "samples/file_input.txt"

	ExecuteFile(inputfile)

}
