// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type AllTripsBatchOption func(*allTripsBatchOptionImpl)

type AllTripsBatchOptions interface {
	Debug() bool
	HasDebug() bool
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	Cursor() string
	HasCursor() bool
	FromTime() time.Time
	HasFromTime() bool
	ToTime() time.Time
	HasToTime() bool
	ToAllTripsOptions() []AllTripsOption
	ToBaseOptions() []BaseOption
	ToTripsOptions() []TripsOption
}

func AllTripsBatchDebug(debug bool) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}
}
func AllTripsBatchDebugFlag(debug *bool) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}
}

func AllTripsBatchSid(sid string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func AllTripsBatchSidFlag(sid *string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func AllTripsBatchCsid(csid string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func AllTripsBatchCsidFlag(csid *string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

func AllTripsBatchCursor(cursor string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}
}
func AllTripsBatchCursorFlag(cursor *string) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}
}

func AllTripsBatchFromTime(fromTime time.Time) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_fromTime = true
		opts.fromTime = fromTime
	}
}
func AllTripsBatchFromTimeFlag(fromTime *time.Time) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if fromTime == nil {
			return
		}
		opts.has_fromTime = true
		opts.fromTime = *fromTime
	}
}

func AllTripsBatchToTime(toTime time.Time) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		opts.has_toTime = true
		opts.toTime = toTime
	}
}
func AllTripsBatchToTimeFlag(toTime *time.Time) AllTripsBatchOption {
	return func(opts *allTripsBatchOptionImpl) {
		if toTime == nil {
			return
		}
		opts.has_toTime = true
		opts.toTime = *toTime
	}
}

type allTripsBatchOptionImpl struct {
	debug        bool
	has_debug    bool
	sid          string
	has_sid      bool
	csid         string
	has_csid     bool
	cursor       string
	has_cursor   bool
	fromTime     time.Time
	has_fromTime bool
	toTime       time.Time
	has_toTime   bool
}

func (a *allTripsBatchOptionImpl) Debug() bool         { return a.debug }
func (a *allTripsBatchOptionImpl) HasDebug() bool      { return a.has_debug }
func (a *allTripsBatchOptionImpl) Sid() string         { return a.sid }
func (a *allTripsBatchOptionImpl) HasSid() bool        { return a.has_sid }
func (a *allTripsBatchOptionImpl) Csid() string        { return a.csid }
func (a *allTripsBatchOptionImpl) HasCsid() bool       { return a.has_csid }
func (a *allTripsBatchOptionImpl) Cursor() string      { return a.cursor }
func (a *allTripsBatchOptionImpl) HasCursor() bool     { return a.has_cursor }
func (a *allTripsBatchOptionImpl) FromTime() time.Time { return a.fromTime }
func (a *allTripsBatchOptionImpl) HasFromTime() bool   { return a.has_fromTime }
func (a *allTripsBatchOptionImpl) ToTime() time.Time   { return a.toTime }
func (a *allTripsBatchOptionImpl) HasToTime() bool     { return a.has_toTime }

type AllTripsBatchParams struct {
	Debug    bool      `json:"debug"`
	Sid      string    `json:"sid"`
	Csid     string    `json:"csid"`
	Cursor   string    `json:"cursor"`
	FromTime time.Time `json:"from_time"`
	ToTime   time.Time `json:"to_time"`
}

func (o AllTripsBatchParams) Options() []AllTripsBatchOption {
	return []AllTripsBatchOption{
		AllTripsBatchDebug(o.Debug),
		AllTripsBatchSid(o.Sid),
		AllTripsBatchCsid(o.Csid),
		AllTripsBatchCursor(o.Cursor),
		AllTripsBatchFromTime(o.FromTime),
		AllTripsBatchToTime(o.ToTime),
	}
}

// ToAllTripsOptions converts AllTripsBatchOption to an array of AllTripsOption
func (o *allTripsBatchOptionImpl) ToAllTripsOptions() []AllTripsOption {
	return []AllTripsOption{
		AllTripsDebug(o.Debug()),
	}
}

// ToBaseOptions converts AllTripsBatchOption to an array of BaseOption
func (o *allTripsBatchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

// ToTripsOptions converts AllTripsBatchOption to an array of TripsOption
func (o *allTripsBatchOptionImpl) ToTripsOptions() []TripsOption {
	return []TripsOption{
		TripsCursor(o.Cursor()),
		TripsFromTime(o.FromTime()),
		TripsToTime(o.ToTime()),
	}
}

func makeAllTripsBatchOptionImpl(opts ...AllTripsBatchOption) *allTripsBatchOptionImpl {
	res := &allTripsBatchOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAllTripsBatchOptions(opts ...AllTripsBatchOption) AllTripsBatchOptions {
	return makeAllTripsBatchOptionImpl(opts...)
}
