package errors

import (
	"errors"
	"fmt"
	"log/slog"
)

var (
	codeUnknown Code = MustNewCode("unknown")
)

// base is the fundamental struct that implements the error interface.
// The order of the struct is 'TCMEUA'.
type base struct {
	// t denotes the custom type of the error.
	t typ
	// c denotes the short code for the error message.
	c Code
	// m contains error message passed through errors.New.
	m string
	// e is the actual error being wrapped.
	e error
	// u denotes the url for the documentation (if present) for the error.
	u string
	// a denotes any additional error messages (if present).
	a []string
}

func (b *base) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("type", b.t.s),
		slog.String("code", b.c.s),
		slog.String("message", b.m),
		slog.String("url", b.u),
		slog.Any("additional", b.a),
	)
}

// base implements Error interface.
func (b *base) Error() string {
	if b.e != nil {
		return b.e.Error()
	}

	return fmt.Sprintf("%s(%s): %s", b.t.s, b.c, b.m)
}

// New returns a base error. It requires type, code and message as input.
func New(t typ, code Code, message string) *base {
	return &base{
		t: t,
		c: code,
		m: message,
		e: nil,
		u: "",
		a: []string{},
	}
}

// Newf returns a new base by formatting the error message with the supplied format specifier.
func Newf(t typ, code Code, format string, args ...interface{}) *base {
	return &base{
		t: t,
		c: code,
		m: fmt.Sprintf(format, args...),
		e: nil,
	}
}

// Wrapf returns a new error by formatting the error message with the supplied format specifier
// and wrapping another error with base.
func Wrapf(cause error, t typ, code Code, format string, args ...interface{}) *base {
	return &base{
		t: t,
		c: code,
		m: fmt.Sprintf(format, args...),
		e: cause,
	}
}

// WithAdditional wraps an existing base error with a new formatted message.
// It is used when the original error already contains type and code.
func WithAdditional(cause error, format string, args ...interface{}) *base {
	t, c, m, e, u, a := Unwrapb(cause)
	b := &base{
		t: t,
		c: c,
		m: m,
		e: e,
		u: u,
		a: a,
	}

	return b.WithAdditional(append(a, fmt.Sprintf(format, args...))...)
}

// WithUrl adds a url to the base error and returns a new base error.
func (b *base) WithUrl(u string) *base {
	return &base{
		t: b.t,
		c: b.c,
		m: b.m,
		e: b.e,
		u: u,
		a: b.a,
	}
}

// WithUrl adds additional messages to the base error and returns a new base error.
func (b *base) WithAdditional(a ...string) *base {
	return &base{
		t: b.t,
		c: b.c,
		m: b.m,
		e: b.e,
		u: b.u,
		a: a,
	}
}

// Unwrapb is a combination of built-in errors.As and type casting.
// It finds the first error in cause that matches base,
// and if one is found, returns the individual fields of base.
// Otherwise, it returns TypeInternal, the original error string
// and the error itself.
//
//lint:ignore ST1008 we want to return arguments in the 'TCMEUA' order of the struct
func Unwrapb(cause error) (typ, Code, string, error, string, []string) {
	base, ok := cause.(*base)
	if ok {
		return base.t, base.c, base.m, base.e, base.u, base.a
	}

	return TypeInternal, codeUnknown, cause.Error(), cause, "", []string{}
}

// Ast checks if the provided error matches the specified custom error type.
func Ast(cause error, typ typ) bool {
	t, _, _, _, _, _ := Unwrapb(cause)

	return t == typ
}

// Ast checks if the provided error matches the specified custom error code.
func Asc(cause error, code Code) bool {
	_, c, _, _, _, _ := Unwrapb(cause)

	return c.s == code.s
}

// Join is a wrapper around errors.Join.
func Join(errs ...error) error {
	return errors.Join(errs...)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func WrapNotFoundf(cause error, code Code, format string, args ...interface{}) *base {
	return Wrapf(cause, TypeNotFound, code, format, args...)
}

func NewNotFoundf(code Code, format string, args ...interface{}) *base {
	return Newf(TypeNotFound, code, format, args...)
}

func WrapInternalf(cause error, code Code, format string, args ...interface{}) *base {
	return Wrapf(cause, TypeInternal, code, format, args...)
}

func NewInternalf(code Code, format string, args ...interface{}) *base {
	return Newf(TypeInternal, code, format, args...)
}

func WrapInvalidInputf(cause error, code Code, format string, args ...interface{}) *base {
	return Wrapf(cause, TypeInvalidInput, code, format, args...)
}

func NewInvalidInputf(code Code, format string, args ...interface{}) *base {
	return Newf(TypeInvalidInput, code, format, args...)
}

func WrapUnexpectedf(cause error, code Code, format string, args ...interface{}) *base {
	return Wrapf(cause, TypeInvalidInput, code, format, args...)
}

func NewUnexpectedf(code Code, format string, args ...interface{}) *base {
	return Newf(TypeInvalidInput, code, format, args...)
}
