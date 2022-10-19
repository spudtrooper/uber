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
	Debug() bool
	HasDebug() bool
	TotalLimit() int
	HasTotalLimit() bool
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
	ToAllTripsOptions() []AllTripsOption
	ToTripsOptions() []TripsOption
	ToBaseOptions() []BaseOption
}

func AllTripsBatchDebug(debug bool) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllTripsBatchDebug(bool %+v)}", debug)}
}
func AllTripsBatchDebugFlag(debug *bool) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllTripsBatchDebug(bool %+v)}", debug)}
}

func AllTripsBatchTotalLimit(totalLimit int) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_totalLimit = true
		opts.totalLimit = totalLimit
	}, fmt.Sprintf("api.AllTripsBatchTotalLimit(int %+v)}", totalLimit)}
}
func AllTripsBatchTotalLimitFlag(totalLimit *int) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if totalLimit == nil {
			return
		}
		opts.has_totalLimit = true
		opts.totalLimit = *totalLimit
	}, fmt.Sprintf("api.AllTripsBatchTotalLimit(int %+v)}", totalLimit)}
}

func AllTripsBatchCursor(cursor string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}, fmt.Sprintf("api.AllTripsBatchCursor(string %+v)}", cursor)}
}
func AllTripsBatchCursorFlag(cursor *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}, fmt.Sprintf("api.AllTripsBatchCursor(string %+v)}", cursor)}
}

func AllTripsBatchFromTime(fromTime time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_fromTime = true
		opts.fromTime = fromTime
	}, fmt.Sprintf("api.AllTripsBatchFromTime(time.Time %+v)}", fromTime)}
}
func AllTripsBatchFromTimeFlag(fromTime *time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if fromTime == nil {
			return
		}
		opts.has_fromTime = true
		opts.fromTime = *fromTime
	}, fmt.Sprintf("api.AllTripsBatchFromTime(time.Time %+v)}", fromTime)}
}

func AllTripsBatchToTime(toTime time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_toTime = true
		opts.toTime = toTime
	}, fmt.Sprintf("api.AllTripsBatchToTime(time.Time %+v)}", toTime)}
}
func AllTripsBatchToTimeFlag(toTime *time.Time) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if toTime == nil {
			return
		}
		opts.has_toTime = true
		opts.toTime = *toTime
	}, fmt.Sprintf("api.AllTripsBatchToTime(time.Time %+v)}", toTime)}
}

func AllTripsBatchSid(sid string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.AllTripsBatchSid(string %+v)}", sid)}
}
func AllTripsBatchSidFlag(sid *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.AllTripsBatchSid(string %+v)}", sid)}
}

func AllTripsBatchCsid(csid string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.AllTripsBatchCsid(string %+v)}", csid)}
}
func AllTripsBatchCsidFlag(csid *string) AllTripsBatchOption {
	return AllTripsBatchOption{func(opts *allTripsBatchOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.AllTripsBatchCsid(string %+v)}", csid)}
}

type allTripsBatchOptionImpl struct {
	debug          bool
	has_debug      bool
	totalLimit     int
	has_totalLimit bool
	cursor         string
	has_cursor     bool
	fromTime       time.Time
	has_fromTime   bool
	toTime         time.Time
	has_toTime     bool
	sid            string
	has_sid        bool
	csid           string
	has_csid       bool
}

func (a *allTripsBatchOptionImpl) Debug() bool         { return a.debug }
func (a *allTripsBatchOptionImpl) HasDebug() bool      { return a.has_debug }
func (a *allTripsBatchOptionImpl) TotalLimit() int     { return a.totalLimit }
func (a *allTripsBatchOptionImpl) HasTotalLimit() bool { return a.has_totalLimit }
func (a *allTripsBatchOptionImpl) Cursor() string      { return a.cursor }
func (a *allTripsBatchOptionImpl) HasCursor() bool     { return a.has_cursor }
func (a *allTripsBatchOptionImpl) FromTime() time.Time { return a.fromTime }
func (a *allTripsBatchOptionImpl) HasFromTime() bool   { return a.has_fromTime }
func (a *allTripsBatchOptionImpl) ToTime() time.Time   { return a.toTime }
func (a *allTripsBatchOptionImpl) HasToTime() bool     { return a.has_toTime }
func (a *allTripsBatchOptionImpl) Sid() string         { return a.sid }
func (a *allTripsBatchOptionImpl) HasSid() bool        { return a.has_sid }
func (a *allTripsBatchOptionImpl) Csid() string        { return a.csid }
func (a *allTripsBatchOptionImpl) HasCsid() bool       { return a.has_csid }

type AllTripsBatchParams struct {
	Debug      bool      `json:"debug"`
	TotalLimit int       `json:"total_limit"`
	Cursor     string    `json:"cursor"`
	FromTime   time.Time `json:"from_time"`
	ToTime     time.Time `json:"to_time"`
	Sid        string    `json:"sid"`
	Csid       string    `json:"csid"`
}

func (o AllTripsBatchParams) Options() []AllTripsBatchOption {
	return []AllTripsBatchOption{
		AllTripsBatchDebug(o.Debug),
		AllTripsBatchTotalLimit(o.TotalLimit),
		AllTripsBatchCursor(o.Cursor),
		AllTripsBatchFromTime(o.FromTime),
		AllTripsBatchToTime(o.ToTime),
		AllTripsBatchSid(o.Sid),
		AllTripsBatchCsid(o.Csid),
	}
}

// ToAllTripsOptions converts AllTripsBatchOption to an array of AllTripsOption
func (o *allTripsBatchOptionImpl) ToAllTripsOptions() []AllTripsOption {
	return []AllTripsOption{
		AllTripsCsid(o.Csid()),
		AllTripsDebug(o.Debug()),
		AllTripsTotalLimit(o.TotalLimit()),
		AllTripsCursor(o.Cursor()),
		AllTripsFromTime(o.FromTime()),
		AllTripsToTime(o.ToTime()),
		AllTripsSid(o.Sid()),
	}
}

// ToTripsOptions converts AllTripsBatchOption to an array of TripsOption
func (o *allTripsBatchOptionImpl) ToTripsOptions() []TripsOption {
	return []TripsOption{
		TripsFromTime(o.FromTime()),
		TripsToTime(o.ToTime()),
		TripsSid(o.Sid()),
		TripsCsid(o.Csid()),
		TripsCursor(o.Cursor()),
	}
}

// ToBaseOptions converts AllTripsBatchOption to an array of BaseOption
func (o *allTripsBatchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
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
