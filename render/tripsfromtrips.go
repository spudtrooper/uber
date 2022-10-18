package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"sort"
	"strings"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
)

//go:embed tmpl/trips.html
var tripsTmpl string

func tripsFromTrips(trips []api.TripsInfoTrip) ([]byte, handler.RendererConfig, error) {
	config := handler.RendererConfig{
		IsFragment: true,
	}

	nonEmpty := func(s string) string {
		if s == "" {
			s = "--"
		}
		return s
	}

	type row struct {
		ID          string
		VehicleName string
		From        string
		To          string
		PickUp      string
		DropOff     string
		JSON        string
		Total       string
		Status      string
	}
	var rows []row
	for _, trip := range trips {
		var jsonObj = struct {
			Data api.TripsInfoTrip
		}{
			Data: trip,
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)
		var from, to string
		if len(trip.Waypoints) >= 2 {
			from = trip.Waypoints[0]
			to = trip.Waypoints[len(trip.Waypoints)-1]
		}
		total := trip.Fare
		if total == "" {
			total = "$0.00"
		}
		rows = append(rows, row{
			ID:          trip.UUID,
			JSON:        jsonStr,
			PickUp:      nonEmpty(strings.TrimSpace(strings.Split(trip.BeginTripTime, "(")[0])),
			DropOff:     nonEmpty(strings.TrimSpace(strings.Split(trip.DropoffTime, "(")[0])),
			From:        nonEmpty(from),
			To:          nonEmpty(to),
			Total:       total,
			VehicleName: trip.VehicleDisplayName,
			Status:      trip.Status,
		})
	}

	sort.Slice(rows, func(i, j int) bool { return rows[i].ID < rows[j].ID })
	var data = struct {
		VehicleViews []row
	}{
		VehicleViews: rows,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, tripsTmpl, "Trips", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
