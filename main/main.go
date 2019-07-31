package main

import (
	"bufio"
	"fmt"
	"github.com/Ivanhahanov/blockchain/test"
	"os"
)

func ReadMessage() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error of input:", err)
	}
	return in.Text()
}

func main() {
	chain := test.InitBlockChain()

	var message, command string
	var exit bool

	for !exit {
		fmt.Printf("enter your command: ")
		command = ReadMessage()
		switch command {
		case "new":
			fmt.Printf("enter message for block: ")
			message = ReadMessage()
			chain.AddBlock(message)
		case "show":
			test.Show_blocks(chain)
		case "exit":
			exit = true
		default:
			fmt.Println("commands: show, new, exit")
		}
	}

}