package render

import (
	"bytes"
	_ "embed"
	"sort"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/uber/api"
)

//go:embed tmpl/status.html
var statusTmpl string

func Status(p any) ([]byte, handler.RendererConfig, error) {
	d := p.(*api.StatusInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}

	vehicleViews, err := d.VehicleViews()
	if err != nil {
		return nil, config, err
	}

	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []api.VehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, statusTmpl, "status", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
