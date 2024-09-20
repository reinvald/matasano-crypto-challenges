package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	decoded, err := decodeHex([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))

	if err != nil {
		fmt.Printf("Error while decoding hex: %s", err)
		return
	}

	base64 := encodeToBase64(decoded)
	fmt.Printf("%s", base64)
}

func decodeHex(encoded []byte) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(encoded)))
	_, err := hex.Decode(decoded, encoded)
	return decoded, err
}

func encodeToBase64(decoded []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(decoded)))
	base64.StdEncoding.Encode(encoded, decoded)
	return encoded
}
