package api

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/bluele/gcache"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

var (
	sid         = flags.String("uber_sid", "uber sid")
	csid        = flags.String("uber_csid", "uber csid")
	userCreds   = flag.String("uber_user_creds", ".user_creds.json", "file with user credentials")
	noUserCreds = flag.Bool("uber_no_user_creds", false, "Don't use user creds event if it exists")
)

// TODO: need to make genopts regex more flexible, need to have "--function Base" exactly here, which isn't great
//go:generate genopts --function Base sid:string csid:string

// Client is a client for uber.com
type Client struct {
	sid, csid string
	urlCache  gcache.Cache
}

// NewClientFromFlags creates a new client from command line flags
func NewClientFromFlags() (*Client, error) {
	if *sid != "" && *csid != "" {
		client := NewClient(*sid, *csid)
		return client, nil
	}
	if *noUserCreds {
		client := NewClient("", "")
		return client, nil
	}
	if *userCreds != "" {
		client, err := NewClientFromFile(*userCreds)
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return nil, errors.Errorf("Must set --user & --token or --creds_file")
}

// NewClient creates a new client directly from the API Key
func NewClient(sid, csid string) *Client {
	return &Client{
		sid:      sid,
		csid:     csid,
		urlCache: gcache.New(20).LRU().Build(),
	}
}

// NewClientFromFile creates a new client from a JSON file like `user_creds-example.json`
func NewClientFromFile(credsFile string) (*Client, error) {
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &Client{
		sid:      creds.SID,
		csid:     creds.CSID,
		urlCache: gcache.New(20).LRU().Build(),
	}, nil
}

type Creds struct {
	SID  string `json:"sid"`
	CSID string `json:"csid"`
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	if !io.FileExists(credsFile) {
		return
	}
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
}

func (c *Client) withAuth(headers map[string]string, optss ...BaseOption) map[string]string {
	opts := MakeBaseOptions(optss...)

	sid := or.String(opts.Sid(), c.sid)
	csid := or.String(opts.Csid(), c.csid)
	headers["cookie"] = request.CreateCookie([][2]string{
		{"sid", sid},
		{"csid", csid},
		{"allow-geolocation", `true`},
	})

	return headers
}
