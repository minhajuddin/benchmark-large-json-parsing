package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(output)), &result)
}
