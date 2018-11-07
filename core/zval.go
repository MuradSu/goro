package core

import (
	"fmt"
	"strconv"
)

type Val interface {
	GetType() ZType
}

type ZVal struct {
	v Val
}

func (z *ZVal) GetType() ZType {
	if z.v == nil {
		return ZtNull
	}
	return z.v.GetType()
}

func (z *ZVal) run(ctx Context) (*ZVal, error) {
	return z, nil
}

func (z *ZVal) IsNull() bool {
	if z == nil {
		return true
	}
	if z.v == nil {
		return true
	}
	return false
}

func (z *ZVal) String() string {
	switch n := z.v.(type) {
	case nil:
		return ""
	case ZBool:
		if n {
			return "1"
		} else {
			return ""
		}
	case ZInt:
		return strconv.FormatInt(int64(n), 10)
	case ZFloat:
		return strconv.FormatFloat(float64(n), 'g', -1, 64)
	case ZString:
		return string(n)
	}
	switch z.GetType() {
	case ZtNull:
		return ""
	case ZtArray:
		return "Array"
	case ZtObject:
		// TODO call __toString()
		return "" // fatal error if no __toString() method
	case ZtResource:
		return "Resource id #" // TODO
	default:
		return fmt.Sprintf("Unknown[%T]", z.v)
	}
}
