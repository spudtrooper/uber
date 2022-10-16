package api

import (
	"encoding/json"
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type getUserInfo struct {
	User struct {
		FirstName                      string `json:"firstName"`
		LastSelectedPaymentProfileUUID string `json:"lastSelectedPaymentProfileUuid"`
		PictureURL                     string `json:"pictureUrl"`
		Tenancy                        string `json:"tenancy"`
		Typename                       string `json:"__typename"`
	} `json:"user"`
	UUID     string `json:"uuid"`
	Typename string `json:"__typename"`
}

//go:generate genopts --params --function GetUser --extends Base
func (c *Client) GetUser(optss ...GetUserOption) (*getUserInfo, error) {
	opts := MakeGetUserOptions(optss...)

	const uri = "https://m.uber.com/graphql"

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	type variables struct {
	}
	type body struct {
		OperationName string    `json:"operationName"`
		Variables     variables `json:"variables"`
		Query         string    `json:"query"`
	}

	b, err := request.JSONMarshal(body{
		OperationName: "GetUser",
		Variables:     variables{},
		Query:         "query GetUser {user {\n    ...UserFragment\n    __typename\n  }\n}\n\nfragment UserFragment on User {\n  user {\n    firstName\n    lastSelectedPaymentProfileUuid\n    pictureUrl\n    tenancy\n    __typename\n  }\n  uuid\n  __typename\n}\n",
	})
	if err != nil {
		return nil, err
	}

	res, err := request.Post(uri, nil, strings.NewReader(string(b)), request.RequestExtraHeaders(headers))
	if err != nil {
		return nil, err
	}
	var payload struct {
		Data struct {
			User getUserInfo `json:"user"`
		} `json:"data"`
	}
	if err := json.Unmarshal(res.Data, &payload); err != nil {
		return nil, err
	}
	return &payload.Data.User, nil

}
