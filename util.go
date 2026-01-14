package app

import (
	"encoding/json"
	"fmt"
	"log"
)

func OutputJSON(v any) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %s", err)
	}
	fmt.Println(string(b))
}
