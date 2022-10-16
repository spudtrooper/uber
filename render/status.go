package render

import (
	"bytes"
	_ "embed"

	"github.com/spudtrooper/uber/api"
)

//go:embed tmpl/status.html
var statusTmpl string

func Status(p any) ([]byte, error) {
	d := p.(*api.StatusInfo)

	// type nearbyVehicle struct {
	// 	Lat float64
	// 	Lng float64
	// }
	// nearbyVehicles []nearbyVehicle
	// for _, v := range d.Eyeball.NearbyVehicles {
	// 	nearbyVehicles = append(nearbyVehicles, nearbyVehicle{
	// 		Lat: v.Vehicle.Latitude,
	// 		Lng: v.Vehicle.Longitude,
	// 	})
	// }
	type vehicleView struct {
		ID              string
		Description     string
		FullDescription string
		ImageURL        string
		Lat, Lng        float64
		IsNearby        bool
	}
	var vehicleViews []vehicleView
	for id, v := range d.City.VehicleViews {
		nb, ok := d.Eyeball.NearbyVehicles[id]
		var lat, lng float64
		var isNearby bool
		if ok {
			lat = nb.Vehicle.Latitude
			lng = nb.Vehicle.Longitude
			isNearby = true
		}
		vehicleViews = append(vehicleViews, vehicleView{
			ID:              id,
			Description:     v.Description,
			FullDescription: v.FullDescription,
			ImageURL:        v.ProductImage.URL,
			Lat:             lat,
			Lng:             lng,
			IsNearby:        isNearby,
		})
	}
	var data = struct {
		// NearbyVehicles []nearbyVehicle
		VehicleViews []vehicleView
	}{
		// NearbyVehicles: nearbyVehicles,
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, statusTmpl, "status", data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
