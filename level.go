package log

import "strconv"

type Level int

const (
	_ Level = iota
	LDebug
	LInfo
	LWarn
	LError
	LPanic
)

func ConvertToLevel(n int) Level {
	if l := Level(n); l <= LPanic && l >= LDebug {
		return l
	}

	return LDebug
}

func (receiver Level) String() string {
	str := ""
	switch receiver {
	case LDebug:
		str = "debug"
	case LInfo:
		str = "info"
	case LWarn:
		str = "warn"
	case LError:
		str = "error"
	case LPanic:
		str = "panic"
	default:
		str = "unknown:" + strconv.Itoa(int(receiver))
	}

	return str
}
