package main

import "testing"

func TestParseArgFlag(t *testing.T) {
	cases := []struct {
		xArg string // execution argument
		eKey string // expected key
		eVal string // expected value
	}{
		{"-key=val", "key", "val"},
		{"-ke-y=val", "ke-y", "val"},
		{"-key=va-l", "key", "va-l"},
	}

	for _, c := range cases {
		gKey, gVal := parseArgFlag(c.xArg)
		if gKey != c.eKey {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eKey, gKey)
		}
		if gVal != c.eVal {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eVal, gVal)
		}
	}
}

func TestAttempConfigOverrideFromArgs(t *testing.T) {
	cases := []struct {
		xConfig config   // config retrieved from local file
		xArgs   []string // config retrieved from execution args (has priority)
		eConfig config   // expected config result
	}{
		{
			config{},
			[]string{},
			config{},
		},
		{
			config{RemoteUrl: "http://asd.com"},
			[]string{},
			config{RemoteUrl: "http://asd.com"},
		},
		{
			config{},
			[]string{"-remoteUrl=http://qwe.com"},
			config{RemoteUrl: "http://qwe.com"},
		},
		{
			config{RemoteUrl: "http://asd.com"},
			[]string{"-remoteUrl=http://qwe.com"},
			config{RemoteUrl: "http://qwe.com"},
		},
	}

	for _, c := range cases {
		attempConfigOverrideFromArgs(&c.xConfig, c.xArgs)
		if c.xConfig != c.eConfig {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eConfig, c.xConfig)
		}
	}
}

func TestVerifyConfigIntegrity(t *testing.T) {
	cases := []struct {
		xConfig config
		eOk     bool
		eErr    string
	}{
		{config{}, false, "[error] remoteUrl cannot be empty"},
		{config{RemoteUrl: "notAUrl"}, false, "[error] remoteUrl is not a valid URL"},
		{config{RemoteUrl: "invalid.url"}, false, "[error] remoteUrl is not a valid URL"},
		{config{RemoteUrl: "http://valid.url"}, true, ""},
		{config{RemoteUrl: "http://valid.url/"}, true, ""},
		{config{RemoteUrl: "http://valid.url/files"}, true, ""},
	}

	for _, c := range cases {
		gOk, gErr := verifyConfigIntegrity(c.xConfig)
		if c.eOk != gOk {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eOk, gOk)
		}
		if c.eErr != gErr {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eErr, gErr)
		}
	}
}
