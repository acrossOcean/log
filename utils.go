package log

import (
	"bufio"
	"bytes"
	"runtime"
	"runtime/debug"
	"strings"
)

func getFileInfo() (fileName string, lineNo int, funcName string) {
	for i := 0; i < 20; i++ {
		pc, file, line, ok := runtime.Caller(i)

		if !ok {
			continue
		}
		if strings.Index(file, "github.com/across!ocean/log") > 0 {
			fileName = file
			lineNo = line
			f := runtime.FuncForPC(pc)
			funcName = f.Name()

			return
		}
	}

	return "", 0, ""
}

func getStackInfo() string {
	var result strings.Builder
	buf := debug.Stack()
	r := bufio.NewReader(bytes.NewReader(buf))
	count := 0
	for {
		line, _, _ := r.ReadLine()
		result.Write(line)
		result.Write([]byte("\n"))
		count++
		if count >= 20 {
			break
		}
	}

	return result.String()
}
