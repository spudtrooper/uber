// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"
)

type AllTripsBatchOption struct {
	f func(*allTripsBatchOptionImpl)
	s string
}

func (o AllTripsBatchOption) String() string { return o.s }

type AllTripsBatchOptions interface {
	Csid() string
	HasCsid() bool
	Cursor() string
	HasCursor() bool
	Debug() bool
	HasDebug() bool
	FromTime() time.Time
	HasFromTime() bool
	Sid() string
	HasSid() bool
	ToTime() time.Time
	HasToTime() bool
	TotalLimit() int
	HasTotalLimit() bool
	ToAllTripsOptions() []AllTripsOption
	ToBaseOptions() []BaseOption
	ToTripsOptions() []TripsOption
}

func AllTripsBatchCsid(csid string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.AllTripsBatchCsid(string %+v)", csid)}
}
func AllTripsBatchCsidFlag(csid *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.AllTripsBatchCsid(string %+v)", csid)}
}

func AllTripsBatchCursor(cursor string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}, fmt.Sprintf("api.AllTripsBatchCursor(string %+v)", cursor)}
}
func AllTripsBatchCursorFlag(cursor *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}, fmt.Sprintf("api.AllTripsBatchCursor(string %+v)", cursor)}
}

func AllTripsBatchDebug(debug bool) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllTripsBatchDebug(bool %+v)", debug)}
}
func AllTripsBatchDebugFlag(debug *bool) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllTripsBatchDebug(bool %+v)", debug)}
}

func AllTripsBatchFromTime(fromTime time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_fromTime = true
		opts.fromTime = fromTime
	}, fmt.Sprintf("api.AllTripsBatchFromTime(time.Time %+v)", fromTime)}
}
func AllTripsBatchFromTimeFlag(fromTime *time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if fromTime == nil {
			return
		}
		opts.has_fromTime = true
		opts.fromTime = *fromTime
	}, fmt.Sprintf("api.AllTripsBatchFromTime(time.Time %+v)", fromTime)}
}

func AllTripsBatchSid(sid string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.AllTripsBatchSid(string %+v)", sid)}
}
func AllTripsBatchSidFlag(sid *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.AllTripsBatchSid(string %+v)", sid)}
}

func AllTripsBatchToTime(toTime time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_toTime = true
		opts.toTime = toTime
	}, fmt.Sprintf("api.AllTripsBatchToTime(time.Time %+v)", toTime)}
}
func AllTripsBatchToTimeFlag(toTime *time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if toTime == nil {
			return
		}
		opts.has_toTime = true
		opts.toTime = *toTime
	}, fmt.Sprintf("api.AllTripsBatchToTime(time.Time %+v)", toTime)}
}

func AllTripsBatchTotalLimit(totalLimit int) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_totalLimit = true
		opts.totalLimit = totalLimit
	}, fmt.Sprintf("api.AllTripsBatchTotalLimit(int %+v)", totalLimit)}
}
func AllTripsBatchTotalLimitFlag(totalLimit *int) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if totalLimit == nil {
			return
		}
		opts.has_totalLimit = true
		opts.totalLimit = *totalLimit
	}, fmt.Sprintf("api.AllTripsBatchTotalLimit(int %+v)", totalLimit)}
}

type allTripsBatchOptionImpl struct {
	csid           string
	has_csid       bool
	cursor         string
	has_cursor     bool
	debug          bool
	has_debug      bool
	fromTime       time.Time
	has_fromTime   bool
	sid            string
	has_sid        bool
	toTime         time.Time
	has_toTime     bool
	totalLimit     int
	has_totalLimit bool
}

func (a *allTripsBatchOptionImpl) Csid() string        { return a.csid }
func (a *allTripsBatchOptionImpl) HasCsid() bool       { return a.has_csid }
func (a *allTripsBatchOptionImpl) Cursor() string      { return a.cursor }
func (a *allTripsBatchOptionImpl) HasCursor() bool     { return a.has_cursor }
func (a *allTripsBatchOptionImpl) Debug() bool         { return a.debug }
func (a *allTripsBatchOptionImpl) HasDebug() bool      { return a.has_debug }
func (a *allTripsBatchOptionImpl) FromTime() time.Time { return a.fromTime }
func (a *allTripsBatchOptionImpl) HasFromTime() bool   { return a.has_fromTime }
func (a *allTripsBatchOptionImpl) Sid() string         { return a.sid }
func (a *allTripsBatchOptionImpl) HasSid() bool        { return a.has_sid }
func (a *allTripsBatchOptionImpl) ToTime() time.Time   { return a.toTime }
func (a *allTripsBatchOptionImpl) HasToTime() bool     { return a.has_toTime }
func (a *allTripsBatchOptionImpl) TotalLimit() int     { return a.totalLimit }
func (a *allTripsBatchOptionImpl) HasTotalLimit() bool { return a.has_totalLimit }

type AllTripsBatchParams struct {
	Csid       string    `json:"csid"`
	Cursor     string    `json:"cursor"`
	Debug      bool      `json:"debug"`
	FromTime   time.Time `json:"from_time"`
	Sid        string    `json:"sid"`
	ToTime     time.Time `json:"to_time"`
	TotalLimit int       `json:"total_limit"`
}

func (o AllTripsBatchParams) Options() []AllTripsBatchOption {
	return []AllTripsBatchOption{
		AllTripsBatchCsid(o.Csid),
		AllTripsBatchCursor(o.Cursor),
		AllTripsBatchDebug(o.Debug),
		AllTripsBatchFromTime(o.FromTime),
		AllTripsBatchSid(o.Sid),
		AllTripsBatchToTime(o.ToTime),
		AllTripsBatchTotalLimit(o.TotalLimit),
	}
}

// ToAllTripsOptions converts AllTripsBatchOption to an array of AllTripsOption
func (o *allTripsBatchOptionImpl) ToAllTripsOptions() []AllTripsOption {
	return []AllTripsOption{
		AllTripsCsid(o.Csid()),
		AllTripsCursor(o.Cursor()),
		AllTripsDebug(o.Debug()),
		AllTripsFromTime(o.FromTime()),
		AllTripsSid(o.Sid()),
		AllTripsToTime(o.ToTime()),
		AllTripsTotalLimit(o.TotalLimit()),
	}
}

// ToBaseOptions converts AllTripsBatchOption to an array of BaseOption
func (o *allTripsBatchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseCsid(o.Csid()),
		BaseSid(o.Sid()),
	}
}

// ToTripsOptions converts AllTripsBatchOption to an array of TripsOption
func (o *allTripsBatchOptionImpl) ToTripsOptions() []TripsOption {
	return []TripsOption{
		TripsCsid(o.Csid()),
		TripsCursor(o.Cursor()),
		TripsFromTime(o.FromTime()),
		TripsSid(o.Sid()),
		TripsToTime(o.ToTime()),
	}
}

func makeAllTripsBatchOptionImpl(opts ...AllTripsBatchOption) *allTripsBatchOptionImpl {
	res := &allTripsBatchOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllTripsBatchOptions(opts ...AllTripsBatchOption) AllTripsBatchOptions {
	return makeAllTripsBatchOptionImpl(opts...)
}
