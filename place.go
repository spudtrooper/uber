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
	uri := request.MakeURL("https://m.uber.com/api/getStatus",
		request.Param{"localeCode", `en`},
	)
	cookie := [][2]string{
		// {"marketing_vistor_id", `5145623a-7713-4213-9a5b-4330c4fff064`},
		// {"segmentCookie", `b`},
		// {"utag_geo_code", `US`},
		// {"_gcl_au", `1.1.1812037979.1664022568`},
		// {"UBER_CONSENTMGR", `1664022585681|consent:true`},
		// {"CONSENTMGR", `c1:1%7Cc2:1%7Cc3:1%7Cc4:1%7Cc5:1%7Cc6:1%7Cc7:1%7Cc8:1%7Cc9:1%7Cc10:1%7Cc11:1%7Cc12:1%7Cc13:1%7Cc14:1%7Cc15:1%7Cts:1664022585681%7Cconsent:true`},
		// {"usl_rollout_id", `258f4f62-8687-4fcf-b929-9a5bb14d6127`},
		// {"udi-id", `vcgecsZHYwbuNu/fBNbUOQX2HTG8ic4z9OGB+fsMpXW1URTvfCz8efEkQ6yAAFL7xYnMwcExFdQ0GV6s+W/P3ZY/2sIOKY/AV9LYaKEWYDGbS/NKHtR1QQ6hGrp9t8DeWbKKls9Ul3OCvLxYJK5xIGwXwQr0+wsyiAQDr3BaeiTwEUOnUe3Z90lmRsLGQ8iEeHoh8ADz4nCdPCM9+CuDoA==lwL64oU5HE6wFkL9I1WRrQ==2jBu1AW3IPwukPcaP/+XzRFCR3Y4Bo5NTiyntBRa/tY=`},
		// {"_gid", `GA1.2.514126880.1665880896`},
		// {"_ua", `{"session_id":"443d9c0d-36ca-4802-bca8-3b2c2dcd1787","session_time_ms":1665880904070}`},
		// {"jwt-session", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InNlc3Npb25fdHlwZSI6ImRlc2t0b3Bfc2Vzc2lvbiIsInRlbmFuY3kiOiJ1YmVyL3Byb2R1Y3Rpb24ifSwiaWF0IjoxNjY1ODgwOTA0LCJleHAiOjE2NjU5NjczMDR9.uev6QlDlTI16FmJdgmf3oCLL8Yw3GtIMRRPh1UWQOKA`},
		// {"allow-geolocation", `true`},
		// {"_ga", `GA1.2.180558629.1664022568`},
		// {"utag_main", `v_id:01836f7953a900075825d7b3bbaa05075001506d00ac8$_sn:4$_se:2$_ss:0$_st:1665924041278$segment:b$optimizely_segment:a$ses_id:1665922239306%3Bexp-session$_pn:1%3Bexp-session`},
		// {"_ga_XTGQLY6KPT", `GS1.1.1665922239.4.1.1665922244.0.0.0`},
		// {"udi-fingerprint", `p6KcpirZEkggaQyYxfMhFbaWdxrfWS5LE9Zn5ptEQJb5jPK4UZ0jQLkLxDxxyiGn4R5NALhpg4H1OvLThreZqg==WFkH+VsIHmmdu1A3/zC83bFL4pVZSPd3O/yH26DxfA8=`},
		// {"_cc", url.QueryEscape(`Aa9RuMTSjK/uzJPgKFwpRS+9`)},

		{"csid", `1.1668472902282.vyd8lD9nfJwioD11SIPpt0uXsdzl10lBU1kxUaAFb8o=`},
		{"sid", `QA.CAESEHZN2t12XEjErN2zmZJjetkYx87bmgYiATEqJGE1M2U3YjEzLWU3ZDUtNGUzOC1hNjkzLTIzZWE5OWQ1NGUxZTI8wtE-vk6V_bi-YtHwVKQ2T3OdWzVFCrn8yWbkeiv2A6xnNqQ_r4xR3TCNThgt7BCG_YbcMdK-sT5AwdorOgExQgh1YmVyLmNvbQ.o8YCgsyXIVWuuGjYuO1X7G0PqLXhuBK0GlLNYEgC_yI`},
	}
	headers := map[string]string{
		"authority":       `m.uber.com`,
		"accept":          `*/*`,
		"accept-language": `en-US,en;q=0.9`,
		"cache-control":   `no-cache`,
		"content-type":    `application/json`,
		// "dnt":                `1`,
		// "origin":             `https://m.uber.com`,
		// "pragma":             `no-cache`,
		// "referer":            `https://m.uber.com/looking`,
		// "sec-ch-ua":          `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		// "sec-ch-ua-mobile":   `?0`,
		// "sec-ch-ua-platform": `"macOS"`,
		// "sec-fetch-dest":     `empty`,
		// "sec-fetch-mode":     `cors`,
		// "sec-fetch-site":     `same-origin`,
		// "user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
		"x-csrf-token": `x`,
		// "x-uber-wa-info":     `TSGFOONJHKJLJHHGGGPMTRLK`,

		"cookie": request.CreateCookie(cookie),
	}
	body := `{"latitude":40.7700742,"longitude":-73.9829869}`

	{
		type Body struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		}

		bodyObject := Body{Latitude: 40.7700742, Longitude: -73.9829869}
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
