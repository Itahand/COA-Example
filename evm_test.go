package main

import (
	"fmt"
	"testing"

	. "github.com/bjartek/overflow/v2"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

// /// TESTING EVM-FLOW COAs /////
func TestCOA(t *testing.T) {

	o, err := OverflowTesting()
	fmt.Print(err)
	assert.NoError(t, err)
	// "Create a COA"
	color.Cyan("Create a COA inside Bob's account")
	o.Tx("create_COA",
		WithSigner("bob"),
	).AssertSuccess(t).Print()
}
