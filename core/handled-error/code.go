package handlederror

import "net/http"

type Code struct {
	code int32
	desc string
}

func NewCode(code int32, desc string) *Code {
	return &Code{
		code: code,
		desc: desc,
	}
}

func (c *Code) Code() int32 {
	return c.code
}

func (c *Code) Desc() string {
	return c.desc
}

var (
	InternalErrorCode = NewCode(http.StatusInternalServerError, "internal error")
)
