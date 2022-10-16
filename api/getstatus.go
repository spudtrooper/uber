package api

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/request"
)

type getStatusInfoVehicleView struct {
	AllowCredits        bool   `json:"allowCredits"`
	Capacity            int    `json:"capacity"`
	Description         string `json:"description"`
	DetailedDescription string `json:"detailedDescription"`
	DisplayName         string `json:"displayName"`
	FullDescription     string `json:"fullDescription"`
	IsCashOnly          bool   `json:"isCashOnly"`
	MapImages           []struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		URL    string `json:"url"`
		Format string `json:"format"`
	} `json:"mapImages"`
	ParentProductTypeUUID string `json:"parentProductTypeUUID"`
	ProductImage          struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		URL    string `json:"url"`
		Format string `json:"format"`
	} `json:"productImage"`
}

type getStatusInfoNearbyVehicle struct {
	EtaString      string `json:"etaString"`
	EtaStringShort string `json:"etaStringShort"`
	Vehicle        struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Course    int     `json:"course"`
		Epoch     int64   `json:"epoch"`
	} `json:"vehicle"`
}

type getStatusInfo struct {
	City struct {
		CityID                     string                              `json:"cityId"`
		DefaultVehicleViewID       int                                 `json:"defaultVehicleViewId"`
		VehicleViews               map[string]getStatusInfoVehicleView `json:"vehicleViews"`
		ProductsUnavailableMessage string                              `json:"productsUnavailableMessage"`
	} `json:"city"`
	ClientStatus struct {
		Status string `json:"status"`
	} `json:"clientStatus"`
	Eyeball struct {
		DynamicFares struct {
		} `json:"dynamicFares"`
		NearbyVehicles map[string]getStatusInfoNearbyVehicle `json:"nearbyVehicles"`
		ReverseGeocode struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"reverseGeocode"`
	} `json:"eyeball"`
	Metadata struct {
		PollingIntervalMs int `json:"pollingIntervalMs"`
	} `json:"metadata"`
	// Trip interface{} `json:"trip"`
}

//go:generate genopts --params --function GetStatus --extends Base localeCode:string:en latitude:float64:40.7701286 longitude:float64:-73.9829762
func (c *Client) GetStatus(optss ...GetStatusOption) (*getStatusInfo, error) {
	opts := MakeGetStatusOptions(optss...)

	uri := request.MakeURL("https://m.uber.com/api/getStatus",
		request.Param{"localeCode", opts.LocaleCode()},
	)

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	type body struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	b, err := request.JSONMarshal(body{
		Latitude:  opts.Latitude(),
		Longitude: opts.Longitude(),
	})
	if err != nil {
		return nil, err
	}

	var payload struct {
		Data   getStatusInfo `json:"data"`
		Status string        `json:"status"`
	}

	if _, err := request.Post(uri, &payload, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}
	if payload.Status != "success" {
		return nil, errors.Errorf("unexpected status: %s", payload.Status)
	}
	return &payload.Data, nil

}
