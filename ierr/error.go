package ierr

import "fmt"

var (
	Success = NewBizErr(0, "Success")

	SystemError = NewBizErr(100000, "SystemError")
)

type IBizError interface {
	Error() string
	Raise()
	Code() int32
	Msg() string
	RaiseIf(assert interface{}, args ...interface{})
}

type BizError struct {
	code int32
	msg  string
}

func NewBizErr(code int32, msg string) IBizError {
	return &BizError{
		code: code,
		msg:  msg,
	}
}

const format = "[BizError] ErrorCode = [%d], Message = [%s]."

func (b *BizError) clone() *BizError {
	return &BizError{
		code: b.code,
		msg:  b.msg,
	}
}

func (b *BizError) raise() {
	panic(b)
}

func (b *BizError) Error() string {
	return fmt.Sprintf(format, b.code, b.msg)
}

func (b *BizError) Raise() {
	err := b.clone()
	err.raise()
}

func (b *BizError) Code() int32 {
	return b.code
}

func (b *BizError) Msg() string {
	return b.msg
}

func (b *BizError) RaiseIf(assert interface{}, args ...interface{}) {
	switch t := assert.(type) {
	case bool:
		if t {
			err := b.clone()
			if len(args) > 0 {
				if s, ok := args[0].(string); ok {
					err.msg = fmt.Sprintf(s, args[1:]...)
				} else {
					err.msg = fmt.Sprint(args)
				}
			}
			err.raise()
		}
	}
}
