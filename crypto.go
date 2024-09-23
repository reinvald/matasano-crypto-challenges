package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {

	/* 1.1 Convert hex to base64 */
	decoded, err := decodeHex([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))

	if err != nil {
		fmt.Printf("Error while decoding hex: %s", err)
		return
	}

	base64 := encodeToBase64(decoded)
	fmt.Printf("%s\n", base64)

	/* 1.2 Fixed XOR */
	decoded1, err := decodeHex([]byte("1c0111001f010100061a024b53535009181c"))
	if err != nil {
		fmt.Printf("Error while decoding hex: %s", err)
		return
	}

	decoded2, err := decodeHex([]byte("686974207468652062756c6c277320657965"))
	if err != nil {
		fmt.Printf("Error while decoding hex: %s", err)
		return
	}

	xorRaw, err := fixedXOR(decoded1, decoded2)
	if err != nil {
		fmt.Printf("Error during XOR: %s", err)
		return
	}

	fmt.Printf("%s", encodeHex(xorRaw))
}

// decodes hex buffer
func decodeHex(encoded []byte) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(encoded)))
	_, err := hex.Decode(decoded, encoded)
	return decoded, err
}

// encodes buffer to hex
func encodeHex(decoded []byte) []byte {
	encoded := make([]byte, hex.EncodedLen(len(decoded)))
	hex.Encode(encoded, decoded)
	return encoded
}

// encodes buffer to base64
func encodeToBase64(decoded []byte) []byte {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(decoded)))
	base64.StdEncoding.Encode(encoded, decoded)
	return encoded
}

// performs bitwise XOR on two buffers of equal length
func fixedXOR(buf1 []byte, buf2 []byte) ([]byte, error) {

	if len(buf1) != len(buf2) {
		return nil, errors.New("Input buffers have different lengths")
	}

	xor := make([]byte, len(buf1))

	for i, v := range buf1 {
		xor[i] = v ^ buf2[i]
	}

	return xor, nil
}
