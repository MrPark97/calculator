package main

// wrongIdError occurs when there is wrong id syntax
type wrongIdError struct {
	s string
}

func (e *wrongIdError) Error() string {
	return e.s
}

// presentation error
type presentationError struct {
	s string
}

func (e *presentationError) Error() string {
	return e.s
}

// division by zero error
type zeroDivisionError struct {
	s string
}

func (e *zeroDivisionError) Error() string {
	return e.s
}

// presentation error
type parenthesesError struct {
	s string
}

func (e *parenthesesError) Error() string {
	return e.s
}
