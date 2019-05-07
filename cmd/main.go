package main

import "Demo-Blockchain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendDate("Send 1 BTC to CCY")
	bc.SendDate("Send 1 BTC to YLJ")
	bc.SendDate("Send 1 BTC to CJW")
	bc.Print()
}
