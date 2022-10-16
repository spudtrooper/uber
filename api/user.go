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

//go:generate genopts --params --function User --extends Base
func (c *Client) User(optss ...UserOption) (*getUserInfo, error) {
	opts := MakeUserOptions(optss...)

	const uri = "https://riders.uber.com/graphql"

	headers := c.withAuth(map[string]string{
		"authority":     `riders.uber.com`,
		"cache-control": `no-cache`,
		"content-type":  `application/json`,
		"x-csrf-token":  `x`,
	}, opts.ToBaseOptions()...)

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
		Query:         "query GetUser {\n  user {\n    ...UserFragment\n    __typename\n  }\n}\n\nfragment UserFragment on User {\n  user {\n    alternateEmails {\n      email\n      paymentProfileUuid\n      __typename\n    }\n    email\n    firstName\n    formattedNumber\n    languageId\n    lastName\n    pictureUrl\n    role\n    signupCountry\n    tenancy\n    __typename\n  }\n  __typename\n}\n",
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
