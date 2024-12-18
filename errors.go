package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("Validation failed")
	ErrSignal = errors.New("Received Signal")
)

// class of errors associated with CI steps
type stepErr struct {
	step  string //record step name
	msg   string //describes error
	cause error  //cause to store the underlying error
}

func (s *stepErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}
 
func (s *stepErr) Is(target error) bool {
	t, ok := target.(*stepErr)
	if !ok {
		return false
	}

	return t.step == s.step
}

func (s *stepErr) Unwrap() error {
	return s.cause
}