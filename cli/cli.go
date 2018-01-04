package cli

import (
	"blockchain/lib"
	"flag"
	"fmt"
	"os"
)

// Cli is used for handling cmd line interface to program
type Cli struct {
	bc *blockchain.BlockChain
}

// NewCli constructs handler for managing cmd line interface
func NewCli(bc *blockchain.BlockChain) *Cli {
	return &Cli{bc}
}

// GetBalance calculates balance amt of all unspent outputs for address
func (cli *Cli) GetBalance(address string) {
	var balance int64

	txOut := cli.bc.GetUnspentOutputs(address)
	for _, output := range txOut {
		balance += output.Value
	}

	fmt.Printf("The Balance of %s is %v", address, balance)
}

// Run starts cmd line interface and parses args
func (cli *Cli) Run() {
	// define two cli input modes
	//addCmdSet := flag.NewFlagSet("add", flag.ExitOnError)
	printCmdSet := flag.NewFlagSet("print", flag.ExitOnError)
	getbalanceCmdSet := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//addCmd := addCmdSet.String("addblock", "", "Text describing transaction.")
	getbalanceCmd := getbalanceCmdSet.String("address", "", "Get the balance of address.")

	if len(os.Args) < 2 {
		fmt.Println("subcommand required.")
		os.Exit(1)
	}

	switch os.Args[1] {
	//case "add":
	//	addCmdSet.Parse(os.Args[2:])
	case "print":
		printCmdSet.Parse(os.Args[2:])
	case "getbalance":
		getbalanceCmdSet.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	//if addCmdSet.Parsed() {
	//	// sanity check
	//	if *addCmd == "" {
	//		addCmdSet.PrintDefaults()
	//		os.Exit(1)
	//	}
	//	cli.bc.AddBlock(*addCmd)
	//}

	if printCmdSet.Parsed() {
		cli.bc.Print()
	}

	if getbalanceCmdSet.Parsed() {
		if *getbalanceCmd == "" {
			getbalanceCmdSet.PrintDefaults()
			os.Exit(1)
		}

		cli.GetBalance(*getbalanceCmd)
	}
}

func (cli *Cli) send(from, to string, amount int) {

	txn := blockchain.NewUserTransaction(from, to, amount, cli.bc)
	cli.bc.MineBlock([]*blockchain.Transaction{txn})
	fmt.Println("success...")
}
