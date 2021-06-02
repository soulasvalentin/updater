package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Updater started")

	args := os.Args[1:]

	// config
	c := config{}
	if err := readConfig(&c); err != nil {
		fmt.Println(err)
	}
	attempConfigOverrideFromArgs(&c, args)
	printConfig(c)
	if cOk, cErr := verifyConfigIntegrity(c); !cOk {
		log.Fatal(cErr)
	}

	// command
	cmd := "sync"
	if len(args) > 0 && args[0] == "build" {
		cmd = "build"
	}

	switch cmd {
	case "sync":
		sync(c)
	case "build":
		build(c)
	}

	fmt.Println("Updater finished")

	if contains(args, "-pause") {
		fmt.Println("")
		fmt.Println("Press any key..")
		os.Stdin.Read([]byte{0})
	}
}
