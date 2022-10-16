package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"sort"

	"github.com/spudtrooper/uber/api"
)

//go:embed tmpl/status.html
var statusTmpl string

func Status(p any) ([]byte, error) {
	d := p.(*api.StatusInfo)

	type vehicleView struct {
		ID                  string
		Description         string
		FullDescription     string
		DetailedDescription string
		ImageURL            string
		Lat, Lng            float64
		IsNearby            bool
		JSON                string
		Name                string
	}
	var vehicleViews []vehicleView
	for id, v := range d.City.VehicleViews {
		nb, ok := d.Eyeball.NearbyVehicles[id]
		var lat, lng float64
		var isNearby bool
		var jsonObj = struct {
			GetStatusInfoNearbyVehicle api.GetStatusInfoNearbyVehicle
			GetStatusInfoVehicleView   api.GetStatusInfoVehicleView
		}{
			GetStatusInfoVehicleView: v,
		}
		if ok {
			lat = nb.Vehicle.Latitude
			lng = nb.Vehicle.Longitude
			isNearby = true
			jsonObj.GetStatusInfoNearbyVehicle = nb
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, err
		}
		jsonStr := string(jsonBytes)
		vehicleViews = append(vehicleViews, vehicleView{
			ID:                  id,
			Description:         v.Description,
			FullDescription:     v.FullDescription,
			DetailedDescription: v.DetailedDescription,
			ImageURL:            v.ProductImage.URL,
			Lat:                 lat,
			Lng:                 lng,
			IsNearby:            isNearby,
			JSON:                jsonStr,
			Name:                v.DisplayName,
		})
	}
	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []vehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, statusTmpl, "status", data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
