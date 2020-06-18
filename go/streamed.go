package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(bufio.NewReader(os.Stdin))
	for dec.More() {
		var result map[string]interface{}
		err := dec.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", result)
	}
}
