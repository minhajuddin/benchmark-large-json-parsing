package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(bufio.NewReader(os.Stdin))

	// read open bracket
	// _, err := dec.Token()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// while the array contains values
	for dec.More() {
		// var m Message
		var result map[string]interface{}
		// decode an array value (Message)
		err := dec.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("%s\n", result)
	}

	// read closing bracket
	// _, err = dec.Token()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
