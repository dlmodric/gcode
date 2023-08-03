// Package gsyscall TODO
package gsyscall

import (
	"fmt"
	"syscall"
)

// Pid TODO
func Pid() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("PID:", pid)
}
