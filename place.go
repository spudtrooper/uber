package main

import (
	"flag"
	"log"
	"strings"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://riders.uber.com/graphql")
	cookie := [][2]string{
		{"marketing_vistor_id", `5145623a-7713-4213-9a5b-4330c4fff064`},
		{"segmentCookie", `b`},
		{"utag_geo_code", `US`},
		{"_gcl_au", `1.1.1812037979.1664022568`},
		{"UBER_CONSENTMGR", `1664022585681|consent:true`},
		{"CONSENTMGR", `c1:1%7Cc2:1%7Cc3:1%7Cc4:1%7Cc5:1%7Cc6:1%7Cc7:1%7Cc8:1%7Cc9:1%7Cc10:1%7Cc11:1%7Cc12:1%7Cc13:1%7Cc14:1%7Cc15:1%7Cts:1664022585681%7Cconsent:true`},
		{"usl_rollout_id", `258f4f62-8687-4fcf-b929-9a5bb14d6127`},
		{"sid", `QA.CAESEHZN2t12XEjErN2zmZJjetkYx87bmgYiATEqJGE1M2U3YjEzLWU3ZDUtNGUzOC1hNjkzLTIzZWE5OWQ1NGUxZTI8wtE-vk6V_bi-YtHwVKQ2T3OdWzVFCrn8yWbkeiv2A6xnNqQ_r4xR3TCNThgt7BCG_YbcMdK-sT5AwdorOgExQgh1YmVyLmNvbQ.o8YCgsyXIVWuuGjYuO1X7G0PqLXhuBK0GlLNYEgC_yI`},
		{"_gid", `GA1.2.514126880.1665880896`},
		{"_ga", `GA1.2.180558629.1664022568`},
		{"utag_main", `v_id:01836f7953a900075825d7b3bbaa05075001506d00ac8$_sn:4$_se:2$_ss:0$_st:1665924041278$segment:b$optimizely_segment:a$ses_id:1665922239306%3Bexp-session$_pn:1%3Bexp-session`},
		{"_ga_XTGQLY6KPT", `GS1.1.1665922239.4.1.1665922244.0.0.0`},
		{"x-uber-analytics-session-id", `57e0dfd8-da30-4c19-8742-cdea005d64cb`},
		{"csid", `1.1668515875704.KeAInu4+GdZPH+y1XoapdTsC0V4AULs1dOG05K8/itE=`},
		{"_ua", `{"session_id":"ff403617-efc8-4596-859a-b7f674ae0696","session_time_ms":1665923875809}`},
		{"jwt-session", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NjU5MjM4NzUsImV4cCI6MTY2NjAxMDI3NX0.2uOLgkW4OH-YG5TE9zRjOX1ZESilUtJ5qMPNGuyohRk`},
		{"optimizelyEndUserId", `oeu1665923875921r0.31827434167484703`},
	}
	headers := map[string]string{
		"authority":          `riders.uber.com`,
		"accept":             `*/*`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"content-type":       `application/json`,
		"dnt":                `1`,
		"origin":             `https://riders.uber.com`,
		"pragma":             `no-cache`,
		"referer":            `https://riders.uber.com/trips?_csid=drPJ2dLBKgCsSVKbKwiA8w&state=wluDzbQy5dF7Nt9sn7mzjmhrdv_vYfMoFgvWpECuIqI%3D&effect=`,
		"sec-ch-ua":          `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-origin`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
		"x-csrf-token":       `x`,

		"cookie": request.CreateCookie(cookie),
	}
	body := `{"operationName":"GetTrips","variables":{"cursor":"","fromTime":null,"toTime":null},"query":"query GetTrips($cursor: String, $fromTime: Float, $toTime: Float) {\n  getTrips(cursor: $cursor, fromTime: $fromTime, toTime: $toTime) {\n    count\n    pagingResult {\n      hasMore\n      nextCursor\n      __typename\n    }\n    reservations {\n      ...TripFragment\n      __typename\n    }\n    trips {\n      ...TripFragment\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TripFragment on Trip {\n  beginTripTime\n  disableCanceling\n  driver\n  dropoffTime\n  fare\n  isRidepoolTrip\n  isScheduledRide\n  isSurgeTrip\n  isUberReserve\n  jobUUID\n  marketplace\n  paymentProfileUUID\n  status\n  uuid\n  vehicleDisplayName\n  waypoints\n  __typename\n}\n"}`

	{
		type Variables struct {
			Cursor   string      `json:"cursor"`
			FromTime interface{} `json:"fromTime"`
			ToTime   interface{} `json:"toTime"`
		}
		type Body struct {
			OperationName string    `json:"operationName"`
			Variables     Variables `json:"variables"`
			Query         string    `json:"query"`
		}

		bodyObject := Body{OperationName: "GetTrips", Variables: Variables{Cursor: "", FromTime: nil, ToTime: nil}, Query: "query GetTrips($cursor: String, $fromTime: Float, $toTime: Float) {\n  getTrips(cursor: $cursor, fromTime: $fromTime, toTime: $toTime) {\n    count\n    pagingResult {\n      hasMore\n      nextCursor\n      __typename\n    }\n    reservations {\n      ...TripFragment\n      __typename\n    }\n    trips {\n      ...TripFragment\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TripFragment on Trip {\n  beginTripTime\n  disableCanceling\n  driver\n  dropoffTime\n  fare\n  isRidepoolTrip\n  isScheduledRide\n  isSurgeTrip\n  isUberReserve\n  jobUUID\n  marketplace\n  paymentProfileUUID\n  status\n  uuid\n  vehicleDisplayName\n  waypoints\n  __typename\n}\n"}
		body = string(request.MustJSONMarshal(bodyObject))
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))

	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", json.MustColorMarshal(payload))
	}
	check.Err(err)
}
