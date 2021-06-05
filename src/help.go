package main

import "fmt"

func printHelp() {
	fmt.Print(`
HELP
===============
The first argument indicates which command will be executed. Available options are:
    > sync (default): Used by end users. Download remote manifest, evaluates which files are missing or changed and downloads originals from public repository (remoteUrl). 
    > build: Generates the manifest file of the current directory (and children). The application itself is excluded from the manifest.`)
}
