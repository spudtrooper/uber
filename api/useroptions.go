// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type UserOption func(*userOptionImpl)

type UserOptions interface {
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
	ToBaseOptions() []BaseOption
}

func UserSid(sid string) UserOption {
	return func(opts *userOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func UserSidFlag(sid *string) UserOption {
	return func(opts *userOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func UserCsid(csid string) UserOption {
	return func(opts *userOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func UserCsidFlag(csid *string) UserOption {
	return func(opts *userOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type userOptionImpl struct {
	sid      string
	has_sid  bool
	csid     string
	has_csid bool
}

func (u *userOptionImpl) Sid() string   { return u.sid }
func (u *userOptionImpl) HasSid() bool  { return u.has_sid }
func (u *userOptionImpl) Csid() string  { return u.csid }
func (u *userOptionImpl) HasCsid() bool { return u.has_csid }

type UserParams struct {
	Sid  string `json:"sid"`
	Csid string `json:"csid"`
}

func (o UserParams) Options() []UserOption {
	return []UserOption{
		UserSid(o.Sid),
		UserCsid(o.Csid),
	}
}

// ToBaseOptions converts UserOption to an array of BaseOption
func (o *userOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseSid(o.Sid()),
		BaseCsid(o.Csid()),
	}
}

func makeUserOptionImpl(opts ...UserOption) *userOptionImpl {
	res := &userOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeUserOptions(opts ...UserOption) UserOptions {
	return makeUserOptionImpl(opts...)
}
