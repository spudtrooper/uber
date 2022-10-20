// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type StatusOption struct {
	f func(*statusOptionImpl)
	s string
}

func (o StatusOption) String() string { return o.s }

type StatusOptions interface {
	Csid() string
	HasCsid() bool
	Latitude() float64
	HasLatitude() bool
	LocaleCode() string
	HasLocaleCode() bool
	Longitude() float64
	HasLongitude() bool
	Sid() string
	HasSid() bool
	ToBaseOptions() []BaseOption
}

func StatusCsid(csid string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.StatusCsid(string %+v)}", csid)}
}
func StatusCsidFlag(csid *string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.StatusCsid(string %+v)}", csid)}
}

func StatusLatitude(latitude float64) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.StatusLatitude(float64 %+v)}", latitude)}
}
func StatusLatitudeFlag(latitude *float64) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.StatusLatitude(float64 %+v)}", latitude)}
}

func StatusLocaleCode(localeCode string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		opts.has_localeCode = true
		opts.localeCode = localeCode
	}, fmt.Sprintf("api.StatusLocaleCode(string %+v)}", localeCode)}
}
func StatusLocaleCodeFlag(localeCode *string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		if localeCode == nil {
			return
		}
		opts.has_localeCode = true
		opts.localeCode = *localeCode
	}, fmt.Sprintf("api.StatusLocaleCode(string %+v)}", localeCode)}
}

func StatusLongitude(longitude float64) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.StatusLongitude(float64 %+v)}", longitude)}
}
func StatusLongitudeFlag(longitude *float64) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.StatusLongitude(float64 %+v)}", longitude)}
}

func StatusSid(sid string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.StatusSid(string %+v)}", sid)}
}
func StatusSidFlag(sid *string) StatusOption {
	return StatusOption{func(opts *statusOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.StatusSid(string %+v)}", sid)}
}

type statusOptionImpl struct {
	csid           string
	has_csid       bool
	latitude       float64
	has_latitude   bool
	localeCode     string
	has_localeCode bool
	longitude      float64
	has_longitude  bool
	sid            string
	has_sid        bool
}

func (s *statusOptionImpl) Csid() string        { return s.csid }
func (s *statusOptionImpl) HasCsid() bool       { return s.has_csid }
func (s *statusOptionImpl) Latitude() float64   { return or.Float64(s.latitude, 40.7701286) }
func (s *statusOptionImpl) HasLatitude() bool   { return s.has_latitude }
func (s *statusOptionImpl) LocaleCode() string  { return or.String(s.localeCode, "en") }
func (s *statusOptionImpl) HasLocaleCode() bool { return s.has_localeCode }
func (s *statusOptionImpl) Longitude() float64  { return or.Float64(s.longitude, -73.9829762) }
func (s *statusOptionImpl) HasLongitude() bool  { return s.has_longitude }
func (s *statusOptionImpl) Sid() string         { return s.sid }
func (s *statusOptionImpl) HasSid() bool        { return s.has_sid }

type StatusParams struct {
	Csid       string  `json:"csid"`
	Latitude   float64 `json:"latitude" default:"40.7701286"`
	LocaleCode string  `json:"locale_code" default:"\"en\""`
	Longitude  float64 `json:"longitude" default:"-73.9829762"`
	Sid        string  `json:"sid"`
}

func (o StatusParams) Options() []StatusOption {
	return []StatusOption{
		StatusCsid(o.Csid),
		StatusLatitude(o.Latitude),
		StatusLocaleCode(o.LocaleCode),
		StatusLongitude(o.Longitude),
		StatusSid(o.Sid),
	}
}

// ToBaseOptions converts StatusOption to an array of BaseOption
func (o *statusOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeStatusOptionImpl(opts ...StatusOption) *statusOptionImpl {
	res := &statusOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeStatusOptions(opts ...StatusOption) StatusOptions {
	return makeStatusOptionImpl(opts...)
}
