package log

import (
	"fmt"
	"time"
)

type Logger struct {
	tags map[string]interface{}

	level     Level
	formatter Formatter
}

var defaultLogger *Logger

func init() {
	defaultLogger = DefaultLogger()
}

func DefaultLogger() *Logger {
	result := &Logger{
		tags: map[string]interface{}{
			"_level": LDebug,
			"_time":  time.Now(),
			"_file":  "",
			"_line":  -1,
			"_func":  "",
		},
		level:     LDebug,
		formatter: jsonFormatter,
	}

	return result
}

func (receiver *Logger) copy() *Logger {
	result := &Logger{
		tags:      map[string]interface{}{},
		formatter: receiver.formatter,
	}

	for k, v := range receiver.tags {
		result.tags[k] = v
	}

	return result
}

func SetStaticTag(k string, v interface{}) {
	defaultLogger.SetStaticTag(k, v)
}

func SetStaticTags(m map[string]interface{}) {
	defaultLogger.SetStaticTags(m)
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

func SetFormatter(formatter Formatter) {
	defaultLogger.SetFormatter(formatter)
}

func WithTag(k string, v interface{}) *Logger {
	return defaultLogger.WithTag(k, v)
}

func WithTags(m map[string]interface{}) *Logger {
	return defaultLogger.WithTags(m)
}

func Debug(first interface{}, more ...interface{}) {
	defaultLogger.Debug(first, more...)
}

func Info(first interface{}, more ...interface{}) {
	defaultLogger.Info(first, more...)
}

func Warn(first interface{}, more ...interface{}) {
	defaultLogger.Warn(first, more...)
}

func Error(first interface{}, more ...interface{}) {
	defaultLogger.Error(first, more...)
}

func Panic(first interface{}, more ...interface{}) {
	defaultLogger.Panic(first, more...)
}

func (receiver *Logger) SetStaticTag(k string, v interface{}) {
	receiver.tags[k] = v
}

func (receiver *Logger) SetStaticTags(m map[string]interface{}) {
	for k, v := range m {
		receiver.tags[k] = v
	}
}

func (receiver *Logger) SetLevel(level Level) {
	receiver.level = level
}

func (receiver *Logger) SetFormatter(formatter Formatter) {
	receiver.formatter = formatter
}

func (receiver *Logger) WithTag(k string, v interface{}) *Logger {
	result := receiver.copy()
	result.tags[k] = v
	return result
}

func (receiver *Logger) WithTags(m map[string]interface{}) *Logger {
	result := receiver.copy()
	for k, v := range m {
		result.tags[k] = v
	}

	return result
}

func (receiver *Logger) Debug(first interface{}, more ...interface{}) {
	var msg string
	switch first := first.(type) {
	case string:
		msg = fmt.Sprintf(first, more...)
	default:
		infos := make([]interface{}, 0)
		infos = append(infos, first)
		infos = append(infos, more...)
		msg = fmt.Sprint(infos...)
	}

	fileName, line, funcName := getFileInfo()
	tags := map[string]interface{}{}
	// 处理tags
	for k, v := range receiver.tags {
		var vv interface{}
		switch k {
		case "_level":
			vv = LDebug.String()
		case "_time":
			vv = time.Now().Format(time.RFC3339)
		case "_file":
			vv = fileName
		case "_line":
			vv = line
		case "_func":
			vv = funcName
		default:
			vv = v
		}

		tags[k] = vv
	}

	fullMsg := receiver.formatter(LDebug, msg, tags)

	receiver.output(Message{msg: fullMsg})
}

func (receiver *Logger) Info(first interface{}, more ...interface{}) {
	if receiver.level > LInfo {
		return
	}

	var msg string
	switch first := first.(type) {
	case string:
		msg = fmt.Sprintf(first, more...)
	default:
		infos := make([]interface{}, 0)
		infos = append(infos, first)
		infos = append(infos, more...)
		msg = fmt.Sprint(infos...)
	}

	fileName, line, funcName := getFileInfo()
	tags := map[string]interface{}{}
	// 处理tags
	for k, v := range receiver.tags {
		var vv interface{}
		switch k {
		case "_level":
			vv = LInfo.String()
		case "_time":
			vv = time.Now().Format(time.RFC3339)
		case "_file":
			vv = fileName
		case "_line":
			vv = line
		case "_func":
			vv = funcName
		default:
			vv = v
		}

		tags[k] = vv
	}

	fullMsg := receiver.formatter(LDebug, msg, tags)

	receiver.output(Message{msg: fullMsg})
}

func (receiver *Logger) Warn(first interface{}, more ...interface{}) {
	if receiver.level > LWarn {
		return
	}

	var msg string
	switch first := first.(type) {
	case string:
		msg = fmt.Sprintf(first, more...)
	default:
		infos := make([]interface{}, 0)
		infos = append(infos, first)
		infos = append(infos, more...)
		msg = fmt.Sprint(infos...)
	}

	fileName, line, funcName := getFileInfo()
	tags := map[string]interface{}{}
	// 处理tags
	for k, v := range receiver.tags {
		var vv interface{}
		switch k {
		case "_level":
			vv = LWarn.String()
		case "_time":
			vv = time.Now().Format(time.RFC3339)
		case "_file":
			vv = fileName
		case "_line":
			vv = line
		case "_func":
			vv = funcName
		default:
			vv = v
		}

		tags[k] = vv
	}

	fullMsg := receiver.formatter(LDebug, msg, tags)

	receiver.output(Message{msg: fullMsg})
}

func (receiver *Logger) Error(first interface{}, more ...interface{}) {
	if receiver.level > LError {
		return
	}

	var msg string
	switch first := first.(type) {
	case string:
		msg = fmt.Sprintf(first, more...)
	default:
		infos := make([]interface{}, 0)
		infos = append(infos, first)
		infos = append(infos, more...)
		msg = fmt.Sprint(infos...)
	}

	fileName, line, funcName := getFileInfo()
	tags := map[string]interface{}{}
	// 处理tags
	for k, v := range receiver.tags {
		var vv interface{}
		switch k {
		case "_level":
			vv = LError.String()
		case "_time":
			vv = time.Now().Format(time.RFC3339)
		case "_file":
			vv = fileName
		case "_line":
			vv = line
		case "_func":
			vv = funcName
		default:
			vv = v
		}

		tags[k] = vv
	}

	fullMsg := receiver.formatter(LDebug, msg, tags)

	receiver.output(Message{msg: fullMsg})
}

func (receiver *Logger) Panic(first interface{}, more ...interface{}) {
	if receiver.level > LPanic {
		return
	}

	var msg string
	switch first := first.(type) {
	case string:
		msg = fmt.Sprintf(first, more...)
	default:
		infos := make([]interface{}, 0)
		infos = append(infos, first)
		infos = append(infos, more...)
		msg = fmt.Sprint(infos...)
	}

	fileName, line, funcName := getFileInfo()
	tags := map[string]interface{}{}
	// 处理tags
	for k, v := range receiver.tags {
		var vv interface{}
		switch k {
		case "_level":
			vv = LPanic.String()
		case "_time":
			vv = time.Now().Format(time.RFC3339)
		case "_file":
			vv = fileName
		case "_line":
			vv = line
		case "_func":
			vv = funcName
		default:
			vv = v
		}

		tags[k] = vv
	}

	// 额外添加 panic 的stack 信息
	tags["_stack"] = getStackInfo()

	fullMsg := receiver.formatter(LDebug, msg, tags)

	receiver.output(Message{msg: fullMsg})
}

func (receiver *Logger) output(msg Message) {
	fmt.Println(msg)
}
