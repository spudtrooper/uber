// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"github.com/spudtrooper/goutil/or"
)

type StatusOption func(*statusOptionImpl)

type StatusOptions interface {
	LocaleCode() string
	HasLocaleCode() bool
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	ToBaseOptions() []BaseOption
}

func StatusLocaleCode(localeCode string) StatusOption {
	return func(opts *statusOptionImpl) {
		opts.has_localeCode = true
		opts.localeCode = localeCode
	}
}
func StatusLocaleCodeFlag(localeCode *string) StatusOption {
	return func(opts *statusOptionImpl) {
		if localeCode == nil {
			return
		}
		opts.has_localeCode = true
		opts.localeCode = *localeCode
	}
}

func StatusLatitude(latitude float64) StatusOption {
	return func(opts *statusOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func StatusLatitudeFlag(latitude *float64) StatusOption {
	return func(opts *statusOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func StatusLongitude(longitude float64) StatusOption {
	return func(opts *statusOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func StatusLongitudeFlag(longitude *float64) StatusOption {
	return func(opts *statusOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func StatusSid(sid string) StatusOption {
	return func(opts *statusOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func StatusSidFlag(sid *string) StatusOption {
	return func(opts *statusOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func StatusCsid(csid string) StatusOption {
	return func(opts *statusOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func StatusCsidFlag(csid *string) StatusOption {
	return func(opts *statusOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type statusOptionImpl struct {
	localeCode     string
	has_localeCode bool
	latitude       float64
	has_latitude   bool
	longitude      float64
	has_longitude  bool
	sid            string
	has_sid        bool
	csid           string
	has_csid       bool
}

func (s *statusOptionImpl) LocaleCode() string  { return or.String(s.localeCode, "en") }
func (s *statusOptionImpl) HasLocaleCode() bool { return s.has_localeCode }
func (s *statusOptionImpl) Latitude() float64   { return or.Float64(s.latitude, 40.7701286) }
func (s *statusOptionImpl) HasLatitude() bool   { return s.has_latitude }
func (s *statusOptionImpl) Longitude() float64  { return or.Float64(s.longitude, -73.9829762) }
func (s *statusOptionImpl) HasLongitude() bool  { return s.has_longitude }
func (s *statusOptionImpl) Sid() string         { return s.sid }
func (s *statusOptionImpl) HasSid() bool        { return s.has_sid }
func (s *statusOptionImpl) Csid() string        { return s.csid }
func (s *statusOptionImpl) HasCsid() bool       { return s.has_csid }

type StatusParams struct {
	LocaleCode string  `json:"locale_code" default:"\"en\""`
	Latitude   float64 `json:"latitude" default:"40.7701286"`
	Longitude  float64 `json:"longitude" default:"-73.9829762"`
	Sid        string  `json:"sid"`
	Csid       string  `json:"csid"`
}

func (o StatusParams) Options() []StatusOption {
	return []StatusOption{
		StatusLocaleCode(o.LocaleCode),
		StatusLatitude(o.Latitude),
		StatusLongitude(o.Longitude),
		StatusSid(o.Sid),
		StatusCsid(o.Csid),
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
		opt(res)
	}
	return res
}

func MakeStatusOptions(opts ...StatusOption) StatusOptions {
	return makeStatusOptionImpl(opts...)
}
