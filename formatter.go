package log

import "encoding/json"

type Formatter func(level Level, msg string, tags map[string]interface{}) string

var jsonFormatter Formatter = func(level Level, msg string, tags map[string]interface{}) string {
	info := map[string]interface{}{
		"level": level,
		"msg":   msg,
	}

	for k, v := range tags {
		info["_"+k] = v
	}

	b, _ := json.Marshal(info)
	return string(b)
}
