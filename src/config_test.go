package main

import (
	"errors"
	"testing"
)

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

func TestVerifyConfigIntegritySync(t *testing.T) {
	cases := []struct {
		xConfig config
		eErr    error
	}{
		{config{}, errors.New("[error] remoteUrl cannot be empty")},
		{config{RemoteUrl: "notAUrl"}, errors.New("[error] remoteUrl is not a valid URL")},
		{config{RemoteUrl: "invalid.url"}, errors.New("[error] remoteUrl is not a valid URL")},
		{config{RemoteUrl: "http://valid.url"}, nil},
		{config{RemoteUrl: "http://valid.url/"}, nil},
		{config{RemoteUrl: "http://valid.url/files"}, nil},
	}

	for _, c := range cases {
		gErr := verifyConfigIntegrity(c.xConfig, "sync")
		if c.eErr == nil {
			if gErr != nil {
				t.Errorf("Failed. Expected 'nil', got '%v'", gErr)
			}
		} else if c.eErr.Error() != gErr.Error() {
			t.Errorf("Failed. Expected '%v', got '%v'", c.eErr, gErr)
		}
	}
}
