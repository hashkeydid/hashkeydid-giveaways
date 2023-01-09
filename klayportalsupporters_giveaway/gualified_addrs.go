package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

var (
	seedHash   string = ""
	totalLucky        = 100
)

func main() {
	var addrs []string

	file, _ := os.Open("klayportalsupporters_giveaway/qualified_addrs.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		addrs = append(addrs, lineText)
	}

	seed := common.HexToHash(seedHash).Big()

	for i := 0; i < totalLucky; i++ {
		luckyIndex := new(big.Int).Mod(seed, big.NewInt(int64(len(addrs)))).Int64()
		fmt.Println(addrs[luckyIndex])
		addrs = append(addrs[0:luckyIndex], addrs[luckyIndex+1:]...)
	}
}
