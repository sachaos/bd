package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PrefixHex = "0x"
const PrefixBin = "0b"

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
	switch origin[:2] {
	case PrefixHex:
		hexStr := strings.TrimPrefix(origin, PrefixHex)
		decoded, err := hex.DecodeString(hexStr)
		if err != nil {
			return nil, err
		}

		return &Binary{
			origin: origin,
			raw:    decoded,
		}, nil
	case PrefixBin:
		var decoded []byte
		var str string
		binStr := strings.TrimPrefix(origin, PrefixBin)

		for i := len(binStr); i > 0; i -= 8 {
			if i - 8 < 0 {
				str = binStr[0:i]
			} else {
				str = binStr[i-8:i]
			}
			v, err := strconv.ParseUint(str, 2, 8)
			if err != nil {
				return nil, err
			}
			decoded = append(decoded, byte(v))
		}

		return &Binary{
			origin: origin,
			raw:    decoded,
		}, nil
	}

	return nil, fmt.Errorf("unknown format")
}

func Describe(b []byte) {
	fmt.Printf("hex: 0x%x\n", b)
	fmt.Printf("string: %s\n", b)
	fmt.Printf("bytes: %v\n", b)
	fmt.Printf("binary: %08b\n", b)
	fmt.Println()
	fmt.Printf("length: %dbit, %dbytes\n", len(b) * 8, len(b))
}
