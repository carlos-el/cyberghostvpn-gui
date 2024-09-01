package commander

import (
	"errors"
	"fmt"
	"strings"
)

var ErrSudoRequired = errors.New("the application requires superuser privileges")

func DetectErrSudoRequiredInMsg(msg string) error {
	if strings.Contains(msg, "without sudo") {
		return ErrSudoRequired
	}
	return nil
}

type ErrCommandParse struct {
	Msg  string
	Text string
	Err  error
}

func (e *ErrCommandParse) Error() string {
	return fmt.Sprintf("%v: err %v, parsing %v", e.Msg, e.Err, e.Text)
}

type ErrCommandSysExecution struct {
	Msg string
	Err error
}

func (e *ErrCommandSysExecution) Error() string {
	return fmt.Sprintf("%v: err %v", e.Msg, e.Err)
}

// CommandError is an error that occurs when a command fails.
// ParseError is an error that occurs when parsing the output of a command.
