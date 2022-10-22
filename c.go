package main

import (
	"flag"
	"log"
	"net/url"
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
	uri := request.MakeURL("https://m.uber.com/graphql")
	cookie := [][2]string{
		{"marketing_vistor_id", `5145623a-7713-4213-9a5b-4330c4fff064`},
		{"segmentCookie", `b`},
		{"utag_geo_code", `US`},
		{"_gcl_au", `1.1.1812037979.1664022568`},
		{"UBER_CONSENTMGR", `1664022585681|consent:true`},
		{"CONSENTMGR", `c1:1%7Cc2:1%7Cc3:1%7Cc4:1%7Cc5:1%7Cc6:1%7Cc7:1%7Cc8:1%7Cc9:1%7Cc10:1%7Cc11:1%7Cc12:1%7Cc13:1%7Cc14:1%7Cc15:1%7Cts:1664022585681%7Cconsent:true`},
		{"usl_rollout_id", `258f4f62-8687-4fcf-b929-9a5bb14d6127`},
		{"udi-id", `vcgecsZHYwbuNu/fBNbUOQX2HTG8ic4z9OGB+fsMpXW1URTvfCz8efEkQ6yAAFL7xYnMwcExFdQ0GV6s+W/P3ZY/2sIOKY/AV9LYaKEWYDGbS/NKHtR1QQ6hGrp9t8DeWbKKls9Ul3OCvLxYJK5xIGwXwQr0+wsyiAQDr3BaeiTwEUOnUe3Z90lmRsLGQ8iEeHoh8ADz4nCdPCM9+CuDoA==lwL64oU5HE6wFkL9I1WRrQ==2jBu1AW3IPwukPcaP/+XzRFCR3Y4Bo5NTiyntBRa/tY=`},
		{"sid", `QA.CAESEHZN2t12XEjErN2zmZJjetkYx87bmgYiATEqJGE1M2U3YjEzLWU3ZDUtNGUzOC1hNjkzLTIzZWE5OWQ1NGUxZTI8wtE-vk6V_bi-YtHwVKQ2T3OdWzVFCrn8yWbkeiv2A6xnNqQ_r4xR3TCNThgt7BCG_YbcMdK-sT5AwdorOgExQgh1YmVyLmNvbQ.o8YCgsyXIVWuuGjYuO1X7G0PqLXhuBK0GlLNYEgC_yI`},
		{"csid", `1.1668472902282.vyd8lD9nfJwioD11SIPpt0uXsdzl10lBU1kxUaAFb8o=`},
		{"_ua", `{"session_id":"443d9c0d-36ca-4802-bca8-3b2c2dcd1787","session_time_ms":1665880904070}`},
		{"allow-geolocation", `true`},
		{"optimizelyEndUserId", `oeu1665923875921r0.31827434167484703`},
		{"_ga", `GA1.2.180558629.1664022568`},
		{"_gid", `GA1.2.1438210382.1666472640`},
		{"utag_main", `v_id:01836f7953a900075825d7b3bbaa05075001506d00ac8$_sn:9$_se:3$_ss:0$_st:1666474442758$segment:b$optimizely_segment:a$ses_id:1666472639712%3Bexp-session$_pn:1%3Bexp-session`},
		{"_ga_XTGQLY6KPT", `GS1.1.1666472640.9.0.1666472644.0.0.0`},
		{"jwt-session", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InNlc3Npb25fdHlwZSI6ImRlc2t0b3Bfc2Vzc2lvbiIsInRlbmFuY3kiOiJ1YmVyL3Byb2R1Y3Rpb24ifSwiaWF0IjoxNjY2NDcyNjQ0LCJleHAiOjE2NjY1NTkwNDR9.89-Ozm3ERgualMTKnAjAM69UnymRZTxSunHCkDYOvqA`},
		{"udi-fingerprint", `wiz6NelqTuW3uWVatlWp1uyZOq4njL5wfjkIN0Memga75bww+wG6e8R9L55xj/YMQV0n9+XoxGyLjT38P/ReLg==SK4tQst8F69NC7NnRMfvjryqu20V7MiJ/1QpwmGGJIs=`},
		{"_cc", url.QueryEscape(`Aa9RuMTSjK/uzJPgKFwpRS+9`)},
	}
	headers := map[string]string{
		"authority":          `m.uber.com`,
		"accept":             `*/*`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"content-type":       `application/json`,
		"dnt":                `1`,
		"origin":             `https://m.uber.com`,
		"pragma":             `no-cache`,
		"referer":            `https://m.uber.com/looking?drop%5B0%5D=%7B%22latitude%22%3A40.7762807%2C%22longitude%22%3A-73.9816219%2C%22addressLine1%22%3A%22A%20Cut%20Above%20Grooming%20Salon%22%2C%22addressLine2%22%3A%22143%20W%2069th%20St%2C%20New%20York%20City%22%2C%22id%22%3A%22f66d0468-56fa-6a0a-5354-7f7677c91bef%22%2C%22provider%22%3A%22uber_places%22%2C%22index%22%3A0%7D&pickup=%7B%22latitude%22%3A40.770237%2C%22longitude%22%3A-73.982789%2C%22addressLine1%22%3A%221865%20Broadway%22%2C%22addressLine2%22%3A%22New%20York%2C%20New%20York%22%2C%22id%22%3A%226527fa52-0fe5-5d44-6865-fa3381af7edb%22%2C%22provider%22%3A%22uber_places%22%2C%22index%22%3A0%7D&vehicle=39`,
		"sec-ch-ua":          `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-origin`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36`,
		"x-csrf-token":       `x`,
		"x-uber-wa-info":     `IHIHPOLMHFNQJHUSCHJKFKNS`,

		"cookie": request.CreateCookie(cookie),
	}

	type PickupLocation struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	type Destination struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	type Variables struct {
		Destination    Destination    `json:"destination"`
		PickupLocation PickupLocation `json:"pickupLocation"`
		VehicleViewIds []int          `json:"vehicleViewIds"`
	}
	type Body struct {
		OperationName string    `json:"operationName"`
		Variables     Variables `json:"variables"`
		Query         string    `json:"query"`
	}

	bodyObject := Body{
		OperationName: "FareEstimate",
		Variables: Variables{
			Destination:    Destination{Latitude: 40.7762807, Longitude: -73.9816219},
			PickupLocation: PickupLocation{Latitude: 40.770237, Longitude: -73.982789},
			VehicleViewIds: []int{},
		},
		Query: "query FareEstimate($pickupLocation: InputLocation\u0021, $destination: InputLocation, $pickupTimeMS: Float, $targetProductType: EnumTargetProductType, $vehicleViewIds: [Int\u0021]\u0021) {\n  fareEstimate(\n    pickupLocation: $pickupLocation\n    destination: $destination\n    pickupTimeMS: $pickupTimeMS\n    targetProductType: $targetProductType\n    vehicleViewIds: $vehicleViewIds\n  ) {\n    ...FareEstimateFragment\n    __typename\n  }\n}\n\nfragment FareEstimateFragment on FareEstimateReturn {\n  fares\n  renderRankingRow\n  vehicleViewsOrder\n  __typename\n}\n",
	}

	body := string(request.MustJSONMarshal(bodyObject))

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers), request.RequestAllowedStatusCodes([]int{200, 400}))

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
