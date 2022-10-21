// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type UserOption struct {
	f func(*userOptionImpl)
	s string
}

func (o UserOption) String() string { return o.s }

type UserOptions interface {
	Csid() string
	HasCsid() bool
	Sid() string
	HasSid() bool
	ToBaseOptions() []BaseOption
}

func UserCsid(csid string) UserOption {
	return UserOption{func(opts *userOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.UserCsid(string %+v)}", csid)}
}
func UserCsidFlag(csid *string) UserOption {
	return UserOption{func(opts *userOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.UserCsid(string %+v)}", csid)}
}

func UserSid(sid string) UserOption {
	return UserOption{func(opts *userOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.UserSid(string %+v)}", sid)}
}
func UserSidFlag(sid *string) UserOption {
	return UserOption{func(opts *userOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.UserSid(string %+v)}", sid)}
}

type userOptionImpl struct {
	csid     string
	has_csid bool
	sid      string
	has_sid  bool
}

func (u *userOptionImpl) Csid() string  { return u.csid }
func (u *userOptionImpl) HasCsid() bool { return u.has_csid }
func (u *userOptionImpl) Sid() string   { return u.sid }
func (u *userOptionImpl) HasSid() bool  { return u.has_sid }

type UserParams struct {
	Csid string `json:"csid"`
	Sid  string `json:"sid"`
}

func (o UserParams) Options() []UserOption {
	return []UserOption{
		UserCsid(o.Csid),
		UserSid(o.Sid),
	}
}

// ToBaseOptions converts UserOption to an array of BaseOption
func (o *userOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseCsid(o.Csid()),
		BaseSid(o.Sid()),
	}
}

func makeUserOptionImpl(opts ...UserOption) *userOptionImpl {
	res := &userOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeUserOptions(opts ...UserOption) UserOptions {
	return makeUserOptionImpl(opts...)
}
