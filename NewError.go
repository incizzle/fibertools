package fibertools

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type RichError struct {
	Stack   []RichStack
	Message string
	Code    int
}

type RichStack struct {
	PCName   string
	Function string
	Line     int
}

func (e *RichError) Error() string {
	return e.Message
}

func (e *RichError) StackTrace() []string {
	var messages []string
	for _, v := range e.Stack {
		messages = append(messages, fmt.Sprintf("[Error] in %s[%s:%d] %s", v.PCName, v.Function, v.Line, e.Message))
	}
	return messages
}

// Error function that can be returned on general errors || panics in order to get stack trace in error handler. Must be used in tandom with fibertools.Recover() for most value.
func NewError(err error) *RichError {
	richErr := RichError{
		Message: err.Error(),
	}

	maxStackAmount := 6
	for i := 1; i < maxStackAmount; i++ {
		pc, fn, line, exists := runtime.Caller(i)
		if exists {
			richErr.Stack = append(richErr.Stack, RichStack{
				PCName:   runtime.FuncForPC(pc).Name(),
				Function: fn,
				Line:     line,
			})
		}
	}
	
	fiberErr, ok := err.(*fiber.Error)
	if !ok {
		richErr.Code = 500
	} else {
		richErr.Code = fiberErr.Code
	}

	return &richErr 
}
