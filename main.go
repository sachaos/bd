package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var output string

func main() {
	flag.StringVar(&output, "output", "", "output format")
	flag.Parse()

	arg := flag.Arg(0)
	parse, err := Parse(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if output == "json" {
		JsonOut(parse)
	} else {
		Describe(parse)
	}
}

func Parse(origin string) ([]byte, error) {
	parsed, err := strconv.ParseInt(origin, 0, 0)
	if err != nil {
		return nil, err
	}

	str := strconv.FormatInt(parsed, 16)
	if len(str)%2 == 1 {
		str = "0" + str
	}

	decodeString, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return decodeString, nil
}

func Describe(b []byte) {
	fmt.Printf("hex: 0x%x\n", b)
	fmt.Printf("string: %s\n", b)
	fmt.Printf("bytes: %v\n", b)
	fmt.Printf("binary: %08b\n", b)
	fmt.Println()
	fmt.Printf("length: %dbit, %dbytes\n", len(b)*8, len(b))
}

func JsonOut(b []byte) {
	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
		"hex":         hex.EncodeToString(b),
		"string":      fmt.Sprintf("%s", b),
		"bytes":       fmt.Sprintf("%v", b),
		"binary":      fmt.Sprintf("%b", b),
		"bit_length":  fmt.Sprintf("%d bit", len(b)*8),
		"byte_length": fmt.Sprintf("%d byte", len(b)),
	})
}
