package errors

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

const DEEP = 2

type errKind int

type trace struct {
	file     string
	line     int
	content  string
	funcName string
}

type CustomError struct {
	err    error
	Kind   reflect.Type
	traces []trace
}

func (e *CustomError) Unwrap() error {
	return e.err
}

func (e *CustomError) Is(target error) bool {
	return target.Error() == e.Error()
}

func (e *CustomError) As(err error, target any) bool {
	_, ok := target.(*CustomError)
	if ok {
		return true
	}
	return false
}

func (e *CustomError) Error() string {
	return e.err.Error()
}

func New(msg string) CustomError {

	if msg == "" {
		return CustomError{err: fmt.Errorf("unknown error")}
	}

	customErr := CustomError{err: fmt.Errorf(msg)}

	skip := DEEP
	for ; skip > 0; skip-- {
		tr := trace{}
		pc, file, line, ok := runtime.Caller(skip)
		if ok {
			lineContent, err := getCodeForLine(file, line)
			if err == nil {
				tr.content = lineContent
			}
			tr.file = file
			tr.line = line
			tr.funcName = runtime.FuncForPC(pc).Name()
			customErr.traces = append(customErr.traces, tr)
		}
	}

	return customErr
}

func StackTrace(e CustomError) string {
	var buf bytes.Buffer

	var next *CustomError

	buf.WriteString(fmt.Sprintf("\n%s", e.err.Error()))

	if e.As(&e, next) {

		stackTrace := e.StackTrace()
		buf.WriteString(fmt.Sprintf("\n%s", stackTrace))
	}

	return buf.String()
}

func (e *CustomError) StackTrace() string {
	var buf bytes.Buffer

	for _, trace := range e.traces {
		if _, err := fmt.Fprintf(&buf, "%s:%d %s content: '%s' \n", trace.file, trace.line, trace.funcName, trace.content); err != nil {
			return ""
		}
	}

	return buf.String()
}

func getCodeForLine(filename string, ln int) (string, error) {

	bf, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2 : ln+2]
	return strings.Replace(lines[1], "\t", "  ", -1), nil
}
