package main

import (
	// Uncomment this line to pass the first stage
	"encoding/json"
	"fmt"
	"os"	
	"github.com/msalbrain/bitorento/pkg/decode"
)


func main() {

	command := os.Args[1]

	if command == "decode" {
		// Uncomment this block to pass the first stage

		bencodedValue := os.Args[2]
		decoded, _, err := decode.DecodeBencode(bencodedValue)
		if err != nil {
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))

	} else if command == "decodef" {
		fileName := os.Args[2]
		
		f, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("had issues reading file")
			return
		}

		decoded, _, err := decode.DecodeBencode(string(f))
		if err != nil {
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))

	}else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
