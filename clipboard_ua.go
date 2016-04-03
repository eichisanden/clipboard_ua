package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/woothee/woothee-go"
	"log"
)

func main() {
	agent, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	result, err := woothee.Parse(agent)
	if err != nil {
		log.Fatalf("Could not parse '%s': %s", agent, err)
	}

	fmt.Printf("Name      : %s\n", result.Name)
	fmt.Printf("Version   : %s\n", result.Version)
	fmt.Printf("Category  : %s\n", result.Category)
	fmt.Printf("Os        : %s\n", result.Os)
	fmt.Printf("OsVersion : %s\n", result.OsVersion)
	fmt.Printf("Vendor    : %s\n", result.Vendor)
	fmt.Printf("Type      : %s\n", result.Type)
}
