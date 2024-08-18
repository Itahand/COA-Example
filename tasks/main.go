package main

import (
	"fmt"
	//if you imports this with .  you do not have to repeat overflow everywhere
	. "github.com/bjartek/overflow/v2"
	"github.com/fatih/color"
)

func main() {

	//start an in memory emulator by default
	o := Overflow(
		WithGlobalPrintOptions(),
	)

	fmt.Println("Testing Contract")
	// fmt.Println("Press any key to continue")
	// fmt.Scanln()

	//
	///// TESTING CADENCE 1.0 /////
	//

	color.Red("Should be able to create a COA")
	// Create COA inside Bob's account
	o.Tx("create_COA",
		WithSigner("bob"),
	)
}
