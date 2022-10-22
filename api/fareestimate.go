package api

import (
	"encoding/json"
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type FareEstimateInfo struct {
	FareEstimate FareEstimateInfoFareEstimate
}

type FareEstimateInfoFareEstimateFare struct {
	AdjustmentValue   string `json:"adjustmentValue"`
	CurrencyCode      string `json:"currencyCode"`
	EstimatedTripTime int    `json:"estimatedTripTime"`
	Etd               struct {
		Meta struct {
			LighthouseRequestUUID string `json:"lighthouseRequestUuid"`
		} `json:"meta"`
		EstimatedTripTime       int    `json:"estimatedTripTime"`
		LegalDisclaimer         string `json:"legalDisclaimer"`
		EstimateRequestTime     int    `json:"estimateRequestTime"`
		EtdDisplayString        string `json:"etdDisplayString"`
		EstimatedSoloOnTripTime int    `json:"estimatedSoloOnTripTime"`
	} `json:"etd"`
	Fare    string `json:"fare"`
	FareRef struct {
		FareFlowUUID    string `json:"fareFlowUUID"`
		FareRequestUUID string `json:"fareRequestUUID"`
		FareSessionUUID string `json:"fareSessionUUID"`
	} `json:"fareRef"`
	FareUUID           string        `json:"fareUUID"`
	HasPromo           bool          `json:"hasPromo"`
	HasRidePass        bool          `json:"hasRidePass"`
	HourlyTiers        []interface{} `json:"hourlyTiers"`
	HourlyOverageRates struct {
	} `json:"hourlyOverageRates"`
	PreAdjustmentValue string `json:"preAdjustmentValue"`
	PricingParams      struct {
		PackageVariantUUID string `json:"packageVariantUUID"`
		FareSessionUUID    string `json:"fareSessionUUID"`
	} `json:"pricingParams"`
	ProductUUID string `json:"productUuid"`
	UpfrontFare struct {
		Capacity       int     `json:"capacity"`
		CurrencyCode   string  `json:"currencyCode"`
		DestinationLat float64 `json:"destinationLat"`
		DestinationLng float64 `json:"destinationLng"`
		Fare           string  `json:"fare"`
		OriginLat      float64 `json:"originLat"`
		OriginLng      float64 `json:"originLng"`
		Signature      struct {
			ExpiresAt int    `json:"expiresAt"`
			IssuedAt  int    `json:"issuedAt"`
			Signature string `json:"signature"`
			Version   string `json:"version"`
		} `json:"signature"`
		VehicleViewID   int    `json:"vehicleViewId"`
		UUID            string `json:"uuid"`
		DynamicFareInfo struct {
			IsSobriety                bool    `json:"isSobriety"`
			Multiplier                float64 `json:"multiplier"`
			UUID                      string  `json:"uuid"`
			SurgeSuppressionThreshold int     `json:"surgeSuppressionThreshold"`
		} `json:"dynamicFareInfo"`
		SurgeMultiplier float64       `json:"surgeMultiplier"`
		ViaLocations    []interface{} `json:"viaLocations"`
	} `json:"upfrontFare"`
}

type FareEstimateInfoFareEstimate struct {
	Fares             map[string]FareEstimateInfoFareEstimateFare `json:"fares"`
	RenderRankingRow  bool                                        `json:"renderRankingRow"`
	VehicleViewsOrder []string                                    `json:"vehicleViewsOrder"`
	Typename          string                                      `json:"__typename"`
}

//go:generate genopts --params --function FareEstimate --extends Base pickupLatitude:float64:40.770237 pickupLongitude:float64:-73.982789  destinationLatitude:float64:40.7762807 destinationLongitude:float64:-73.9816219
func (c *Client) FareEstimate(optss ...FareEstimateOption) (*FareEstimateInfo, error) {
	opts := MakeFareEstimateOptions(optss...)

	const uri = "https://m.uber.com/graphql"

	headers := c.withAuth(map[string]string{
		"authority":     `m.uber.com`,
		"cache-control": `no-cache`,
		"content-type":  `application/json`,
		"x-csrf-token":  `x`,
	}, opts.ToBaseOptions()...)

	type point struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	type variables struct {
		Destination    point `json:"destination"`
		PickupLocation point `json:"pickupLocation"`
		VehicleViewIds []int `json:"vehicleViewIds"`
	}
	type body struct {
		OperationName string    `json:"operationName"`
		Variables     variables `json:"variables"`
		Query         string    `json:"query"`
	}

	b, err := request.JSONMarshal(body{
		OperationName: "FareEstimate",
		Variables: variables{
			Destination: point{
				Latitude:  opts.DestinationLatitude(),
				Longitude: opts.DestinationLongitude(),
			},
			PickupLocation: point{
				Latitude:  opts.PickupLatitude(),
				Longitude: opts.PickupLongitude(),
			},
			VehicleViewIds: []int{},
		},
		Query: "query FareEstimate($pickupLocation: InputLocation!, $destination: InputLocation, $pickupTimeMS: Float, $targetProductType: EnumTargetProductType, $vehicleViewIds: [Int!]!) {\n  fareEstimate(\n    pickupLocation: $pickupLocation\n    destination: $destination\n    pickupTimeMS: $pickupTimeMS\n    targetProductType: $targetProductType\n    vehicleViewIds: $vehicleViewIds\n  ) {\n    ...FareEstimateFragment\n    __typename\n  }\n}\n\nfragment FareEstimateFragment on FareEstimateReturn {\n  fares\n  renderRankingRow\n  vehicleViewsOrder\n  __typename\n}\n",
	})
	if err != nil {
		return nil, err
	}

	resp, err := request.Post(uri, nil, strings.NewReader(string(b)), request.RequestExtraHeaders(headers))
	if err != nil {
		return nil, err
	}

	var payload = struct {
		Data struct {
			FareEstimate FareEstimateInfoFareEstimate `json:"fareEstimate"`
		} `json:"data"`
	}{}
	if err := json.Unmarshal(resp.Data, &payload); err != nil {
		return nil, err
	}
	res := &FareEstimateInfo{
		FareEstimate: payload.Data.FareEstimate,
	}
	return res, nil

}
