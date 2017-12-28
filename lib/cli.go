package blockchain

import (
	"flag"
	"fmt"
	"os"
)

// Cli is used for handling cmd line interface to program
type Cli struct {
	bc *BlockChain
}

// NewCli constructs handler for managing cmd line interface
func NewCli(bc *BlockChain) *Cli {
	return &Cli{bc}
}

// Run starts cmd line interface and parses args
func (cli *Cli) Run() {
	// define two cli input modes
	addCmdSet := flag.NewFlagSet("add", flag.ExitOnError)
	printCmdSet := flag.NewFlagSet("print", flag.ExitOnError)

	addCmd := addCmdSet.String("addblock", "", "Text describing transaction")

	if len(os.Args) < 2 {
		fmt.Println("add or print subcommand required.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmdSet.Parse(os.Args[2:])
	case "print":
		printCmdSet.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if addCmdSet.Parsed() {
		// sanity check
		if *addCmd == "" {
			addCmdSet.PrintDefaults()
			os.Exit(1)
		}
		cli.bc.AddBlock(*addCmd)
	}

	if printCmdSet.Parsed() {
		cli.bc.Print()
	}
}
