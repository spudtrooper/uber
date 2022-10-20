// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type BaseOption struct {
	f func(*baseOptionImpl)
	s string
}

func (o BaseOption) String() string { return o.s }

type BaseOptions interface {
	Csid() string
	HasCsid() bool
	Sid() string
	HasSid() bool
}

func BaseCsid(csid string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}, fmt.Sprintf("api.BaseCsid(string %+v)}", csid)}
}
func BaseCsidFlag(csid *string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}, fmt.Sprintf("api.BaseCsid(string %+v)}", csid)}
}

func BaseSid(sid string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}, fmt.Sprintf("api.BaseSid(string %+v)}", sid)}
}
func BaseSidFlag(sid *string) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}, fmt.Sprintf("api.BaseSid(string %+v)}", sid)}
}

type baseOptionImpl struct {
	sid      string
	has_sid  bool
	csid     string
	has_csid bool
}

func (b *baseOptionImpl) Csid() string  { return b.csid }
func (b *baseOptionImpl) HasCsid() bool { return b.has_csid }
func (b *baseOptionImpl) Sid() string   { return b.sid }
func (b *baseOptionImpl) HasSid() bool  { return b.has_sid }

func makeBaseOptionImpl(opts ...BaseOption) *baseOptionImpl {
	res := &baseOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeBaseOptions(opts ...BaseOption) BaseOptions {
	return makeBaseOptionImpl(opts...)
}
