// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type GetUserOption func(*getUserOptionImpl)

type GetUserOptions interface {
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	ToBaseOptions() []BaseOption
}

func GetUserSid(sid string) GetUserOption {
	return func(opts *getUserOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func GetUserSidFlag(sid *string) GetUserOption {
	return func(opts *getUserOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func GetUserCsid(csid string) GetUserOption {
	return func(opts *getUserOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func GetUserCsidFlag(csid *string) GetUserOption {
	return func(opts *getUserOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type getUserOptionImpl struct {
	sid      string
	has_sid  bool
	csid     string
	has_csid bool
}

func (g *getUserOptionImpl) Sid() string   { return g.sid }
func (g *getUserOptionImpl) HasSid() bool  { return g.has_sid }
func (g *getUserOptionImpl) Csid() string  { return g.csid }
func (g *getUserOptionImpl) HasCsid() bool { return g.has_csid }

type GetUserParams struct {
	Sid  string `json:"sid"`
	Csid string `json:"csid"`
}

func (o GetUserParams) Options() []GetUserOption {
	return []GetUserOption{
		GetUserSid(o.Sid),
		GetUserCsid(o.Csid),
	}
}

// ToBaseOptions converts GetUserOption to an array of BaseOption
func (o *getUserOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeGetUserOptionImpl(opts ...GetUserOption) *getUserOptionImpl {
	res := &getUserOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGetUserOptions(opts ...GetUserOption) GetUserOptions {
	return makeGetUserOptionImpl(opts...)
}
