package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const PrefixHex = "0x"

func main() {
	arg := os.Args[1]
	parse, err := Parse(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	Describe(parse.raw)
}

type Binary struct {
	origin string
	raw []byte
}

func Parse(origin string) (*Binary, error) {
	if strings.HasPrefix(origin, PrefixHex) {
		hexStr := strings.TrimPrefix(origin, PrefixHex)
		decoded, err := hex.DecodeString(hexStr)
		if err != nil {
			return nil, err
		}

		return &Binary{
			origin: origin,
			raw:    decoded,
		}, nil
	}

	return nil, nil
}

func Describe(b []byte) {
	fmt.Printf("hex: 0x%x\n", b)
	fmt.Printf("string: %s\n", b)
	fmt.Printf("bytes: %v\n", b)
	fmt.Printf("binary: %08b\n", b)
	fmt.Println()
	fmt.Printf("length: %dbit, %dbytes\n", len(b) * 8, len(b))
}
