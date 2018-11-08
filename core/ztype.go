package core

type ZType int

const (
	ZtNull ZType = iota
	ZtBool
	ZtInt
	ZtFloat
	ZtString
	ZtArray
	ZtObject
	ZtResource
)

// scalar stuff
type ZBool bool
type ZInt int64
type ZFloat float64
type ZString string

func (z ZBool) GetType() ZType {
	return ZtBool
}

func (z ZString) GetType() ZType {
	return ZtString
}

func (z ZInt) GetType() ZType {
	return ZtInt
}

func (z ZFloat) GetType() ZType {
	return ZtFloat
}

func ZStr(s string) *ZVal {
	return &ZVal{ZString(s)}
}
