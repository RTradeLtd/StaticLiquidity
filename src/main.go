package main

// used to test the contract
import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/RTradeLtd/StaticLiquidity/src/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	tokenContract = common.HexToAddress("0x675b45856257CeEf650100C7Ca1b2E8c6FF42e7C")
)

func main() {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), password)
	if err != nil {
		log.Fatal(err)
	}
	_, tx, contractHandler, err := bindings.DeployStaticLiquidity(auth, client, tokenContract)
	if err != nil {
		log.Fatalf("failed to deploy contract %v", err)
	}

	contractAddress, err := bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("failed to wait for deployment tx to be mined %v", err)
	}

	if contractAddress == common.HexToAddress("") {
		log.Fatal("zero address returned for contract address")
	}

	admin, err := contractHandler.Admin(nil)
	if err != nil {
		log.Fatal(err)
	}
	if admin != auth.From {
		log.Fatal("invalid admin address detected")
	}
	fmt.Println("Contract address is ", contractAddress.String())
}
