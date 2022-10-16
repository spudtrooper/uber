// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type TripsOption func(*tripsOptionImpl)

type TripsOptions interface {
	Cursor() string
	HasCursor() bool
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
	cursor     string
	has_cursor bool
	sid        string
	has_sid    bool
	csid       string
	has_csid   bool
}

func (t *tripsOptionImpl) Cursor() string  { return t.cursor }
func (t *tripsOptionImpl) HasCursor() bool { return t.has_cursor }
func (t *tripsOptionImpl) Sid() string     { return t.sid }
func (t *tripsOptionImpl) HasSid() bool    { return t.has_sid }
func (t *tripsOptionImpl) Csid() string    { return t.csid }
func (t *tripsOptionImpl) HasCsid() bool   { return t.has_csid }

type TripsParams struct {
	Cursor string `json:"cursor"`
	Sid    string `json:"sid"`
	Csid   string `json:"csid"`
}

func (o TripsParams) Options() []TripsOption {
	return []TripsOption{
		TripsCursor(o.Cursor),
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
