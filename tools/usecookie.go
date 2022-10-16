package main

import (
	"flag"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/uber/api"
)

func readTokens(cookie string) (sid, csid string, err error) {
	for _, part := range strings.Split(cookie, "; ") {
		parts := strings.SplitN(part, "=", 2)
		if len(parts) != 2 {
			err = errors.Errorf("malformed cookie part: %q", part)
			return
		}
		key, val := parts[0], parts[1]
		if key == "sid" {
			sid = val
		}
		if key == "csid" {
			csid = val
		}
	}
	return
}

func useCookie() error {
	cookie := flag.Arg(0)
	if cookie == "" {
		return errors.Errorf("cookie should be the first arg")
	}
	sid, csid, err := readTokens(cookie)
	if err != nil {
		return err
	}
	if sid == "" {
		return errors.Errorf("no sid")
	}
	if csid == "" {
		return errors.Errorf("no csid")
	}
	creds, err := api.ReadCredsFromFlags()
	if err != nil {
		return errors.Errorf("api.ReadCredsFromFlags: %v", err)
	}
	creds.SID = sid
	creds.CSID = csid
	if err := api.WriteCredsFromFlags(creds); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(useCookie())
}
