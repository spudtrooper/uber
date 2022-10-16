// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"github.com/spudtrooper/goutil/or"
)

type GetStatusOption func(*getStatusOptionImpl)

type GetStatusOptions interface {
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

func GetStatusLocaleCode(localeCode string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		opts.has_localeCode = true
		opts.localeCode = localeCode
	}
}
func GetStatusLocaleCodeFlag(localeCode *string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		if localeCode == nil {
			return
		}
		opts.has_localeCode = true
		opts.localeCode = *localeCode
	}
}

func GetStatusLatitude(latitude float64) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func GetStatusLatitudeFlag(latitude *float64) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func GetStatusLongitude(longitude float64) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func GetStatusLongitudeFlag(longitude *float64) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func GetStatusSid(sid string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func GetStatusSidFlag(sid *string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func GetStatusCsid(csid string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func GetStatusCsidFlag(csid *string) GetStatusOption {
	return func(opts *getStatusOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type getStatusOptionImpl struct {
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

func (g *getStatusOptionImpl) LocaleCode() string  { return or.String(g.localeCode, "en") }
func (g *getStatusOptionImpl) HasLocaleCode() bool { return g.has_localeCode }
func (g *getStatusOptionImpl) Latitude() float64   { return or.Float64(g.latitude, 40.7701286) }
func (g *getStatusOptionImpl) HasLatitude() bool   { return g.has_latitude }
func (g *getStatusOptionImpl) Longitude() float64  { return or.Float64(g.longitude, -73.9829762) }
func (g *getStatusOptionImpl) HasLongitude() bool  { return g.has_longitude }
func (g *getStatusOptionImpl) Sid() string         { return g.sid }
func (g *getStatusOptionImpl) HasSid() bool        { return g.has_sid }
func (g *getStatusOptionImpl) Csid() string        { return g.csid }
func (g *getStatusOptionImpl) HasCsid() bool       { return g.has_csid }

type GetStatusParams struct {
	LocaleCode string  `json:"locale_code" default:"\"en\""`
	Latitude   float64 `json:"latitude" default:"40.7701286"`
	Longitude  float64 `json:"longitude" default:"-73.9829762"`
	Sid        string  `json:"sid"`
	Csid       string  `json:"csid"`
}

func (o GetStatusParams) Options() []GetStatusOption {
	return []GetStatusOption{
		GetStatusLocaleCode(o.LocaleCode),
		GetStatusLatitude(o.Latitude),
		GetStatusLongitude(o.Longitude),
		GetStatusSid(o.Sid),
		GetStatusCsid(o.Csid),
	}
}

// ToBaseOptions converts GetStatusOption to an array of BaseOption
func (o *getStatusOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeGetStatusOptionImpl(opts ...GetStatusOption) *getStatusOptionImpl {
	res := &getStatusOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetStatusOptions(opts ...GetStatusOption) GetStatusOptions {
	return makeGetStatusOptionImpl(opts...)
}
