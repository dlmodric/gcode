// Package main _
package main

import (
	"fmt"

	"github.com/dlmodric/gcode.git/add"
)

func main() {
	res := add.AddV2(3, 4)
	fmt.Println(res)
}
