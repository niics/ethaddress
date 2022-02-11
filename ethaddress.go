package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

func main() {
	// input public key
	pubKey := "040dc6785daab4f31d390cf53a9c04a39088380ac9d4828da48f5e90273d51bb0ceaac56d963074282d87907d98487ba9c29e5719409c4cc42c8d646ae7c190b42"
	pubKeyByte := []byte(pubKey)
	fmt.Println("PubKey: ", pubKey)
	fmt.Println("PubKeyByte: ", pubKeyByte)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKeyByte[1:])
	fmt.Println("address: ", hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
