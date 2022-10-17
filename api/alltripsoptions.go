// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type AllTripsOption struct {
	f func(*allTripsOptionImpl)
	s string
}

func (o AllTripsOption) String() string { return o.s }

type AllTripsOptions interface {
	Debug() bool
	HasDebug() bool
	Cursor() string
	HasCursor() bool
	FromTime() time.Time
	HasFromTime() bool
	ToTime() time.Time
	HasToTime() bool
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	ToTripsOptions() []TripsOption
	ToBaseOptions() []BaseOption
}

func AllTripsDebug(debug bool) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllTripsDebug(bool %+v)}", debug)}
}
func AllTripsDebugFlag(debug *bool) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllTripsDebug(bool %+v)}", debug)}
}

func AllTripsCursor(cursor string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}, fmt.Sprintf("api.AllTripsCursor(string %+v)}", cursor)}
}
func AllTripsCursorFlag(cursor *string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}, fmt.Sprintf("api.AllTripsCursor(string %+v)}", cursor)}
}

func AllTripsFromTime(fromTime time.Time) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_fromTime = true
		opts.fromTime = fromTime
	}, fmt.Sprintf("api.AllTripsFromTime(time.Time %+v)}", fromTime)}
}
func AllTripsFromTimeFlag(fromTime *time.Time) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if fromTime == nil {
			return
		}
		opts.has_fromTime = true
		opts.fromTime = *fromTime
	}, fmt.Sprintf("api.AllTripsFromTime(time.Time %+v)}", fromTime)}
}

func AllTripsToTime(toTime time.Time) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_toTime = true
		opts.toTime = toTime
	}, fmt.Sprintf("api.AllTripsToTime(time.Time %+v)}", toTime)}
}
func AllTripsToTimeFlag(toTime *time.Time) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if toTime == nil {
			return
		}
		opts.has_toTime = true
		opts.toTime = *toTime
	}, fmt.Sprintf("api.AllTripsToTime(time.Time %+v)}", toTime)}
}

func AllTripsSid(sid string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.AllTripsSid(string %+v)}", sid)}
}
func AllTripsSidFlag(sid *string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.AllTripsSid(string %+v)}", sid)}
}

func AllTripsCsid(csid string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.AllTripsCsid(string %+v)}", csid)}
}
func AllTripsCsidFlag(csid *string) AllTripsOption {
	return AllTripsOption{func(opts *allTripsOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.AllTripsCsid(string %+v)}", csid)}
}

type allTripsOptionImpl struct {
	debug        bool
	has_debug    bool
	cursor       string
	has_cursor   bool
	fromTime     time.Time
	has_fromTime bool
	toTime       time.Time
	has_toTime   bool
	sid          string
	has_sid      bool
	csid         string
	has_csid     bool
}

func (a *allTripsOptionImpl) Debug() bool         { return a.debug }
func (a *allTripsOptionImpl) HasDebug() bool      { return a.has_debug }
func (a *allTripsOptionImpl) Cursor() string      { return a.cursor }
func (a *allTripsOptionImpl) HasCursor() bool     { return a.has_cursor }
func (a *allTripsOptionImpl) FromTime() time.Time { return a.fromTime }
func (a *allTripsOptionImpl) HasFromTime() bool   { return a.has_fromTime }
func (a *allTripsOptionImpl) ToTime() time.Time   { return a.toTime }
func (a *allTripsOptionImpl) HasToTime() bool     { return a.has_toTime }
func (a *allTripsOptionImpl) Sid() string         { return a.sid }
func (a *allTripsOptionImpl) HasSid() bool        { return a.has_sid }
func (a *allTripsOptionImpl) Csid() string        { return a.csid }
func (a *allTripsOptionImpl) HasCsid() bool       { return a.has_csid }

type AllTripsParams struct {
	Debug    bool      `json:"debug"`
	Cursor   string    `json:"cursor"`
	FromTime time.Time `json:"from_time"`
	ToTime   time.Time `json:"to_time"`
	Sid      string    `json:"sid"`
	Csid     string    `json:"csid"`
}

func (o AllTripsParams) Options() []AllTripsOption {
	return []AllTripsOption{
		AllTripsDebug(o.Debug),
		AllTripsCursor(o.Cursor),
		AllTripsFromTime(o.FromTime),
		AllTripsToTime(o.ToTime),
		AllTripsSid(o.Sid),
		AllTripsCsid(o.Csid),
	}
}

// ToTripsOptions converts AllTripsOption to an array of TripsOption
func (o *allTripsOptionImpl) ToTripsOptions() []TripsOption {
	return []TripsOption{
		TripsCursor(o.Cursor()),
		TripsFromTime(o.FromTime()),
		TripsToTime(o.ToTime()),
		TripsSid(o.Sid()),
		TripsCsid(o.Csid()),
	}
}

// ToBaseOptions converts AllTripsOption to an array of BaseOption
func (o *allTripsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeAllTripsOptionImpl(opts ...AllTripsOption) *allTripsOptionImpl {
	res := &allTripsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllTripsOptions(opts ...AllTripsOption) AllTripsOptions {
	return makeAllTripsOptionImpl(opts...)
}
