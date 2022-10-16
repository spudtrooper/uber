package api

import (
	"encoding/json"
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type getTripsInfoTrip struct {
	BeginTripTime      string   `json:"beginTripTime"`
	DisableCanceling   bool     `json:"disableCanceling"`
	Driver             string   `json:"driver"`
	DropoffTime        string   `json:"dropoffTime"`
	Fare               string   `json:"fare"`
	IsRidepoolTrip     bool     `json:"isRidepoolTrip"`
	IsScheduledRide    bool     `json:"isScheduledRide"`
	IsSurgeTrip        bool     `json:"isSurgeTrip"`
	IsUberReserve      bool     `json:"isUberReserve"`
	JobUUID            string   `json:"jobUUID"`
	Marketplace        string   `json:"marketplace"`
	PaymentProfileUUID string   `json:"paymentProfileUUID"`
	Status             string   `json:"status"`
	UUID               string   `json:"uuid"`
	VehicleDisplayName string   `json:"vehicleDisplayName"`
	Waypoints          []string `json:"waypoints"`
	Typename           string   `json:"__typename"`
}

type getTripsInfo struct {
	Reservations []interface{}      `json:"reservations"`
	Trips        []getTripsInfoTrip `json:"trips"`
	Typename     string             `json:"__typename"`
	PagingResult struct {
		HasMore    bool   `json:"hasMore"`
		NextCursor string `json:"nextCursor"`
		Typename   string `json:"__typename"`
	} `json:"pagingResult"`
	Count int `json:"count"`
}

//go:generate genopts --params --function Trips --extends Base cursor:string fromTime:time.Time toTime:time.Time
func (c *Client) Trips(optss ...TripsOption) (*getTripsInfo, error) {
	opts := MakeTripsOptions(optss...)

	const uri = "https://riders.uber.com/graphql"

	headers := c.withAuth(map[string]string{
		"authority":     `riders.uber.com`,
		"cache-control": `no-cache`,
		"content-type":  `application/json`,
		"x-csrf-token":  `x`,
	}, opts.ToBaseOptions()...)

	type variables struct {
		Cursor   string `json:"cursor"`
		FromTime int64  `json:"fromTime"`
		ToTime   int64  `json:"toTime"`
	}
	type body struct {
		OperationName string    `json:"operationName"`
		Variables     variables `json:"variables"`
		Query         string    `json:"query"`
	}

	vars := variables{
		Cursor: opts.Cursor(),
	}
	if !opts.FromTime().IsZero() {
		vars.FromTime = opts.FromTime().UnixMilli()
	}
	if !opts.ToTime().IsZero() {
		vars.ToTime = opts.ToTime().UnixMilli()
	}

	b, err := request.JSONMarshal(body{
		OperationName: "GetTrips",
		Variables:     vars,
		Query:         "query GetTrips($cursor: String, $fromTime: Float, $toTime: Float) {\n  getTrips(cursor: $cursor, fromTime: $fromTime, toTime: $toTime) {\n    count\n    pagingResult {\n      hasMore\n      nextCursor\n      __typename\n    }\n    reservations {\n      ...TripFragment\n      __typename\n    }\n    trips {\n      ...TripFragment\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TripFragment on Trip {\n  beginTripTime\n  disableCanceling\n  driver\n  dropoffTime\n  fare\n  isRidepoolTrip\n  isScheduledRide\n  isSurgeTrip\n  isUberReserve\n  jobUUID\n  marketplace\n  paymentProfileUUID\n  status\n  uuid\n  vehicleDisplayName\n  waypoints\n  __typename\n}\n",
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
			GetTrips getTripsInfo `json:"getTrips"`
		} `json:"data"`
	}
	if err := json.Unmarshal(res.Data, &payload); err != nil {
		return nil, err
	}
	return &payload.Data.GetTrips, nil

}
