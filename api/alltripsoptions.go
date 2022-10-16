// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type AllTripsOption func(*allTripsOptionImpl)

type AllTripsOptions interface {
	Debug() bool
	HasDebug() bool
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	Cursor() string
	HasCursor() bool
	ToBaseOptions() []BaseOption
	ToTripsOptions() []TripsOption
}

func AllTripsDebug(debug bool) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}
}
func AllTripsDebugFlag(debug *bool) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}
}

func AllTripsSid(sid string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func AllTripsSidFlag(sid *string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func AllTripsCsid(csid string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func AllTripsCsidFlag(csid *string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

func AllTripsCursor(cursor string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		opts.has_cursor = true
		opts.cursor = cursor
	}
}
func AllTripsCursorFlag(cursor *string) AllTripsOption {
	return func(opts *allTripsOptionImpl) {
		if cursor == nil {
			return
		}
		opts.has_cursor = true
		opts.cursor = *cursor
	}
}

type allTripsOptionImpl struct {
	debug      bool
	has_debug  bool
	sid        string
	has_sid    bool
	csid       string
	has_csid   bool
	cursor     string
	has_cursor bool
}

func (a *allTripsOptionImpl) Debug() bool     { return a.debug }
func (a *allTripsOptionImpl) HasDebug() bool  { return a.has_debug }
func (a *allTripsOptionImpl) Sid() string     { return a.sid }
func (a *allTripsOptionImpl) HasSid() bool    { return a.has_sid }
func (a *allTripsOptionImpl) Csid() string    { return a.csid }
func (a *allTripsOptionImpl) HasCsid() bool   { return a.has_csid }
func (a *allTripsOptionImpl) Cursor() string  { return a.cursor }
func (a *allTripsOptionImpl) HasCursor() bool { return a.has_cursor }

type AllTripsParams struct {
	Debug  bool   `json:"debug"`
	Sid    string `json:"sid"`
	Csid   string `json:"csid"`
	Cursor string `json:"cursor"`
}

func (o AllTripsParams) Options() []AllTripsOption {
	return []AllTripsOption{
		AllTripsDebug(o.Debug),
		AllTripsSid(o.Sid),
		AllTripsCsid(o.Csid),
		AllTripsCursor(o.Cursor),
	}
}

// ToBaseOptions converts AllTripsOption to an array of BaseOption
func (o *allTripsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

// ToTripsOptions converts AllTripsOption to an array of TripsOption
func (o *allTripsOptionImpl) ToTripsOptions() []TripsOption {
	return []TripsOption{
		TripsCursor(o.Cursor()),
	}
}

func makeAllTripsOptionImpl(opts ...AllTripsOption) *allTripsOptionImpl {
	res := &allTripsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAllTripsOptions(opts ...AllTripsOption) AllTripsOptions {
	return makeAllTripsOptionImpl(opts...)
}
