package gsyscall

import (
	"fmt"
	"testing"

	"github.com/dlmodric/gcode.git/add"
)

func TestPid(t *testing.T) {
	Pid()
}

func TestAdd(t *testing.T) {
	res := add.Add(71, 12)
	fmt.Println(res)
}
