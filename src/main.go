package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Updater started")

	args := os.Args[1:]

	// command
	cmd := DefaultCommand
	if len(args) > 0 {
		cmd = args[0]
	}
	fmt.Println("Command:", cmd)

	// config
	c := config{}
	if err := readConfig(&c); err != nil {
		fmt.Println(err)
	}
	attempConfigOverrideFromArgs(&c, args)
	printConfig(c)
	if err := verifyConfigIntegrity(c, cmd); err != nil {
		log.Fatal(err)
	}

	// exec
	switch cmd {
	case "sync":
		sync(c)
	case "build":
		build(c)
	case "help", "-h", "--help":
		printHelp()
	default:
		log.Fatal("Unknown command " + cmd)
	}

	fmt.Println("Updater finished")

	if contains(args, "-pause") {
		fmt.Println("")
		fmt.Println("Press any key..")
		os.Stdin.Read([]byte{0})
	}
}
