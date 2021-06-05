package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

// Read config from local file and returns typed object.
// If file is missing a default is created and exits.
// If json parse fails, exits.
func readConfig(c *config) error {
	fmt.Println("Reading local config..")

	// read config file from local file
	jsonFile, err := os.Open(ConfigFileName)
	if err != nil {
		return errors.New("Config file missing or failed to load. " + err.Error())
	}
	defer jsonFile.Close() // closes the stream at the end of the function

	// transform payload into type
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if errUnmarshal := json.Unmarshal(byteValue, &c); errUnmarshal != nil {
		return errors.New("Error while parsing config json file" + errUnmarshal.Error())
	}

	setDefaultValues(c)

	return nil
}

func setDefaultValues(c *config) {
	if len(c.LocalDirectory) == 0 {
		c.LocalDirectory = DefaultLocalDirectory
	}
}

func printConfig(c config) {
	fmt.Println(" - remoteUrl: " + c.RemoteUrl)
	fmt.Println(" - localDirectory: " + c.LocalDirectory)
}

// Verifies config fields integrity, returns error
func verifyConfigIntegrity(c config, executionMode string) error {
	if len(c.LocalDirectory) == 0 {
		return errors.New("[error] LocalDirectory cannot be empty")
	}
	if executionMode == "sync" {
		if len(c.RemoteUrl) == 0 {
			return errors.New("[error] remoteUrl cannot be empty")
		}
		if _, err := url.ParseRequestURI(c.RemoteUrl); err != nil {
			return errors.New("[error] remoteUrl is not a valid URL")
		}
	}
	return nil
}

// Gets execution arguments and attempts to override config settings.
// Expected format: "-key=val".
func attempConfigOverrideFromArgs(c *config, args []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") && strings.Contains(arg, "=") {
			key, val := parseArgFlag(arg)
			// todo: replace switch with reflec, duh
			switch key {
			case "remoteUrl":
				c.RemoteUrl = val
			case "localDirectory":
				c.LocalDirectory = val
			}
		}
	}
}

// Parse "-key=val" format.
// Assumes chars '-' and '=' are present
func parseArgFlag(arg string) (key string, val string) {
	arg = strings.TrimPrefix(arg, "-")
	s := strings.Split(arg, "=")
	return s[0], s[1]
}
