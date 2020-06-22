package log

type Level int

const (
	LDebug Level = iota
	LInfo
	LWarn
	LError
	LPanic
)
