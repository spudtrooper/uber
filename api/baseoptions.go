// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type BaseOption func(*baseOptionImpl)

type BaseOptions interface {
	Sid() string
	HasSid() bool
	Csid() string
	HasCsid() bool
}

func BaseSid(sid string) BaseOption {
	return func(opts *baseOptionImpl) {
		opts.has_sid = true
		opts.sid = sid
	}
}
func BaseSidFlag(sid *string) BaseOption {
	return func(opts *baseOptionImpl) {
		if sid == nil {
			return
		}
		opts.has_sid = true
		opts.sid = *sid
	}
}

func BaseCsid(csid string) BaseOption {
	return func(opts *baseOptionImpl) {
		opts.has_csid = true
		opts.csid = csid
	}
}
func BaseCsidFlag(csid *string) BaseOption {
	return func(opts *baseOptionImpl) {
		if csid == nil {
			return
		}
		opts.has_csid = true
		opts.csid = *csid
	}
}

type baseOptionImpl struct {
	sid      string
	has_sid  bool
	csid     string
	has_csid bool
}

func (b *baseOptionImpl) Sid() string   { return b.sid }
func (b *baseOptionImpl) HasSid() bool  { return b.has_sid }
func (b *baseOptionImpl) Csid() string  { return b.csid }
func (b *baseOptionImpl) HasCsid() bool { return b.has_csid }

func makeBaseOptionImpl(opts ...BaseOption) *baseOptionImpl {
	res := &baseOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeBaseOptions(opts ...BaseOption) BaseOptions {
	return makeBaseOptionImpl(opts...)
}
