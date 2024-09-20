package main

import (
	"fmt"
	"io/ioutil"
	"log"

	//if you imports this with .  you do not have to repeat overflow everywhere
	. "github.com/bjartek/overflow/v2"
	"github.com/fatih/color"
)

func readJSFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {

	// Specify the path to your JavaScript file
	filePath := "bytecode/GameItem.js"
	// Read the content of the JavaScript file
	jsContent, err := readJSFile(filePath)
	if err != nil {
		log.Fatalf("Error reading JavaScript file: %v", err)
	}
	fmt.Println(jsContent[0])
	//start an in memory emulator by default
	o := Overflow(
		WithGlobalPrintOptions(),
		WithNetwork("testnet"),
	)

	fmt.Println("Testing Solidity interactions with Cadence")
	// fmt.Println("Press any key to continue")
	// fmt.Scanln()

	//
	///// TESTING CADENCE 1.0 /////
	//

	color.Red("Should be able to create a COA")
	// Create COA inside Bob's account
	/* 	o.Tx("create_COA",
		WithSigner("gamer"),
	).Print() */

	/* 	o.Script("get_COA_address",
		WithArg("address", "gamer"),
		WithArg("pathId", 0),
	).Print() */

	// Deploy a Solidity contract to the COA
	color.Cyan("Deploy a Solidity contract to Random's COA")
	/* 	o.Tx("deploy_sol_contract",
		WithSigner("gamer"),
		WithArg("code", jsContent),
		WithArg("pathId", 0),
	).Print() */

	color.Cyan("Mint a game item from the Solidity contract")
	o.Script("call_sol_function",
		WithArg("hexEncodedAddress", "b93cB988D0722E17B67A5E169a47FB6F3A4dea1b"),
		WithArg("address", "gamer"),
		WithArg("pathId", 0),
	).Print()

}
