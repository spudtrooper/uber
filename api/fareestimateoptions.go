// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type FareEstimateOption struct {
	f func(*fareEstimateOptionImpl)
	s string
}

func (o FareEstimateOption) String() string { return o.s }

type FareEstimateOptions interface {
	Csid() string
	HasCsid() bool
	DestinationLatitude() float64
	HasDestinationLatitude() bool
	DestinationLongitude() float64
	HasDestinationLongitude() bool
	PickupLatitude() float64
	HasPickupLatitude() bool
	PickupLongitude() float64
	HasPickupLongitude() bool
	Sid() string
	HasSid() bool
	ToBaseOptions() []BaseOption
}

func FareEstimateCsid(csid string) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.FareEstimateCsid(string %+v)", csid)}
}
func FareEstimateCsidFlag(csid *string) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.FareEstimateCsid(string %+v)", csid)}
}

func FareEstimateDestinationLatitude(destinationLatitude float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_destinationLatitude = true
		opts.destinationLatitude = destinationLatitude
	}, fmt.Sprintf("api.FareEstimateDestinationLatitude(float64 %+v)", destinationLatitude)}
}
func FareEstimateDestinationLatitudeFlag(destinationLatitude *float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if destinationLatitude == nil {
			return
		}
		opts.has_destinationLatitude = true
		opts.destinationLatitude = *destinationLatitude
	}, fmt.Sprintf("api.FareEstimateDestinationLatitude(float64 %+v)", destinationLatitude)}
}

func FareEstimateDestinationLongitude(destinationLongitude float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_destinationLongitude = true
		opts.destinationLongitude = destinationLongitude
	}, fmt.Sprintf("api.FareEstimateDestinationLongitude(float64 %+v)", destinationLongitude)}
}
func FareEstimateDestinationLongitudeFlag(destinationLongitude *float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if destinationLongitude == nil {
			return
		}
		opts.has_destinationLongitude = true
		opts.destinationLongitude = *destinationLongitude
	}, fmt.Sprintf("api.FareEstimateDestinationLongitude(float64 %+v)", destinationLongitude)}
}

func FareEstimatePickupLatitude(pickupLatitude float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_pickupLatitude = true
		opts.pickupLatitude = pickupLatitude
	}, fmt.Sprintf("api.FareEstimatePickupLatitude(float64 %+v)", pickupLatitude)}
}
func FareEstimatePickupLatitudeFlag(pickupLatitude *float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if pickupLatitude == nil {
			return
		}
		opts.has_pickupLatitude = true
		opts.pickupLatitude = *pickupLatitude
	}, fmt.Sprintf("api.FareEstimatePickupLatitude(float64 %+v)", pickupLatitude)}
}

func FareEstimatePickupLongitude(pickupLongitude float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_pickupLongitude = true
		opts.pickupLongitude = pickupLongitude
	}, fmt.Sprintf("api.FareEstimatePickupLongitude(float64 %+v)", pickupLongitude)}
}
func FareEstimatePickupLongitudeFlag(pickupLongitude *float64) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if pickupLongitude == nil {
			return
		}
		opts.has_pickupLongitude = true
		opts.pickupLongitude = *pickupLongitude
	}, fmt.Sprintf("api.FareEstimatePickupLongitude(float64 %+v)", pickupLongitude)}
}

func FareEstimateSid(sid string) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.FareEstimateSid(string %+v)", sid)}
}
func FareEstimateSidFlag(sid *string) FareEstimateOption {
	return FareEstimateOption{func(opts *fareEstimateOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.FareEstimateSid(string %+v)", sid)}
}

type fareEstimateOptionImpl struct {
	csid                     string
	has_csid                 bool
	destinationLatitude      float64
	has_destinationLatitude  bool
	destinationLongitude     float64
	has_destinationLongitude bool
	pickupLatitude           float64
	has_pickupLatitude       bool
	pickupLongitude          float64
	has_pickupLongitude      bool
	sid                      string
	has_sid                  bool
}

func (f *fareEstimateOptionImpl) Csid() string  { return f.csid }
func (f *fareEstimateOptionImpl) HasCsid() bool { return f.has_csid }
func (f *fareEstimateOptionImpl) DestinationLatitude() float64 {
	return or.Float64(f.destinationLatitude, 40.7762807)
}
func (f *fareEstimateOptionImpl) HasDestinationLatitude() bool { return f.has_destinationLatitude }
func (f *fareEstimateOptionImpl) DestinationLongitude() float64 {
	return or.Float64(f.destinationLongitude, -73.9816219)
}
func (f *fareEstimateOptionImpl) HasDestinationLongitude() bool { return f.has_destinationLongitude }
func (f *fareEstimateOptionImpl) PickupLatitude() float64 {
	return or.Float64(f.pickupLatitude, 40.770237)
}
func (f *fareEstimateOptionImpl) HasPickupLatitude() bool { return f.has_pickupLatitude }
func (f *fareEstimateOptionImpl) PickupLongitude() float64 {
	return or.Float64(f.pickupLongitude, -73.982789)
}
func (f *fareEstimateOptionImpl) HasPickupLongitude() bool { return f.has_pickupLongitude }
func (f *fareEstimateOptionImpl) Sid() string              { return f.sid }
func (f *fareEstimateOptionImpl) HasSid() bool             { return f.has_sid }

type FareEstimateParams struct {
	Csid                 string  `json:"csid"`
	DestinationLatitude  float64 `json:"destination_latitude" default:"40.7762807"`
	DestinationLongitude float64 `json:"destination_longitude" default:"-73.9816219"`
	PickupLatitude       float64 `json:"pickup_latitude" default:"40.770237"`
	PickupLongitude      float64 `json:"pickup_longitude" default:"-73.982789"`
	Sid                  string  `json:"sid"`
}

func (o FareEstimateParams) Options() []FareEstimateOption {
	return []FareEstimateOption{
		FareEstimateCsid(o.Csid),
		FareEstimateDestinationLatitude(o.DestinationLatitude),
		FareEstimateDestinationLongitude(o.DestinationLongitude),
		FareEstimatePickupLatitude(o.PickupLatitude),
		FareEstimatePickupLongitude(o.PickupLongitude),
		FareEstimateSid(o.Sid),
	}
}

// ToBaseOptions converts FareEstimateOption to an array of BaseOption
func (o *fareEstimateOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseCsid(o.Csid()),
		BaseSid(o.Sid()),
	}
}

func makeFareEstimateOptionImpl(opts ...FareEstimateOption) *fareEstimateOptionImpl {
	res := &fareEstimateOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFareEstimateOptions(opts ...FareEstimateOption) FareEstimateOptions {
	return makeFareEstimateOptionImpl(opts...)
}
