// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "time"

type TripsOption func(*tripsOptionImpl)

type TripsOptions interface {
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
	ToBaseOptions() []BaseOption
}

func TripsCursor(cursor string) TripsOption {
	return func(opts *tripsOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}
}
func TripsCursorFlag(cursor *string) TripsOption {
	return func(opts *tripsOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}
}

func TripsFromTime(fromTime time.Time) TripsOption {
	return func(opts *tripsOptionImpl) {
		opts.has_fromTime = true
		opts.fromTime = fromTime
	}
}
func TripsFromTimeFlag(fromTime *time.Time) TripsOption {
	return func(opts *tripsOptionImpl) {
		if fromTime == nil {
			return
		}
		opts.has_fromTime = true
		opts.fromTime = *fromTime
	}
}

func TripsToTime(toTime time.Time) TripsOption {
	return func(opts *tripsOptionImpl) {
		opts.has_toTime = true
		opts.toTime = toTime
	}
}
func TripsToTimeFlag(toTime *time.Time) TripsOption {
	return func(opts *tripsOptionImpl) {
		if toTime == nil {
			return
		}
		opts.has_toTime = true
		opts.toTime = *toTime
	}
}

func TripsSid(sid string) TripsOption {
	return func(opts *tripsOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func TripsSidFlag(sid *string) TripsOption {
	return func(opts *tripsOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func TripsCsid(csid string) TripsOption {
	return func(opts *tripsOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func TripsCsidFlag(csid *string) TripsOption {
	return func(opts *tripsOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type tripsOptionImpl struct {
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

func (t *tripsOptionImpl) Cursor() string      { return t.cursor }
func (t *tripsOptionImpl) HasCursor() bool     { return t.has_cursor }
func (t *tripsOptionImpl) FromTime() time.Time { return t.fromTime }
func (t *tripsOptionImpl) HasFromTime() bool   { return t.has_fromTime }
func (t *tripsOptionImpl) ToTime() time.Time   { return t.toTime }
func (t *tripsOptionImpl) HasToTime() bool     { return t.has_toTime }
func (t *tripsOptionImpl) Sid() string         { return t.sid }
func (t *tripsOptionImpl) HasSid() bool        { return t.has_sid }
func (t *tripsOptionImpl) Csid() string        { return t.csid }
func (t *tripsOptionImpl) HasCsid() bool       { return t.has_csid }

type TripsParams struct {
	Cursor   string    `json:"cursor"`
	FromTime time.Time `json:"from_time"`
	ToTime   time.Time `json:"to_time"`
	Sid      string    `json:"sid"`
	Csid     string    `json:"csid"`
}

func (o TripsParams) Options() []TripsOption {
	return []TripsOption{
		TripsCursor(o.Cursor),
		TripsFromTime(o.FromTime),
		TripsToTime(o.ToTime),
		TripsSid(o.Sid),
		TripsCsid(o.Csid),
	}
}

// ToBaseOptions converts TripsOption to an array of BaseOption
func (o *tripsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeTripsOptionImpl(opts ...TripsOption) *tripsOptionImpl {
	res := &tripsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeTripsOptions(opts ...TripsOption) TripsOptions {
	return makeTripsOptionImpl(opts...)
}
