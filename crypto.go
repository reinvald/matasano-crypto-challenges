package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	decodeHex([]byte("48656c6c6f20476f7068657221"))
}

func decodeHex(encoded []byte) ([]byte, error) {

	decoded := make([]byte, hex.DecodedLen(len(encoded)))
	numBytes, err := hex.Decode(decoded, encoded)

	fmt.Printf("%s\n", decoded[:numBytes])

	return decoded, err
}
