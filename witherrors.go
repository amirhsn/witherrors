package witherrors

import "errors"

type withErrors struct {
	code          string
	customMessage interface{}
	dependency    string
	err           error
	message       string
	priority      string
}

type WithErrors interface {
	error
	// Clear existing error with all custom `with` attributes.
	Clear()
	// Clear existing error with all custom `with` attributes and get the struct pointer
	// to be passed into new witherrors struct methods.
	ClearAndGet() WithErrors
	// Get error code.
	GetCode() string
	// Get additional custom error message.
	GetCustomMessage() interface{}
	// Get error dependency.
	GetDependency() string
	// Get error
	GetError() error
	// Get additional error message.
	GetMessage() string
	// Get error priority.
	GetPriority() string
	// Set error code.
	WithCode(c string) WithErrors
	// Set additional custom error message.
	WithCustomMessage(m interface{}) WithErrors
	// Set error dependency.
	WithDependency(d string) WithErrors
	// Set error priority.
	WithPriority(p Priority) WithErrors
}

// Create new error.
// The behaviour is same as go built in error.
func NewError(m string) error {
	if m == "" {
		return errors.New(emptyErrStr)
	}
	return errors.New(m)
}

// Wrap existing error to be passed into available struct attributes.
func WrapError(err error) WithErrors {
	return &withErrors{
		err: err,
	}
}

// Get the error message, implementing go built in Error() function.
func (w *withErrors) Error() string {
	return w.err.Error()
}

func (w *withErrors) Clear() {
	w.code = ""
	w.customMessage = nil
	w.dependency = ""
	w.err = nil
	w.message = ""
	w.priority = ""
}

func (w *withErrors) ClearAndGet() WithErrors {
	w.code = ""
	w.customMessage = nil
	w.dependency = ""
	w.err = nil
	w.message = ""
	w.priority = ""
	return w
}

func (w *withErrors) GetCode() string {
	if w.code == "" {
		w.message = emptyCodeStr
	}
	return w.code
}

func (w *withErrors) GetCustomMessage() interface{} {
	if w.customMessage == nil {
		w.message = emptyCustomMessageStr
	}
	return w.customMessage
}

func (w *withErrors) GetDependency() string {
	if w.dependency == "" {
		w.message = emptyDependencyStr
	}
	return w.dependency
}

func (w *withErrors) GetError() error {
	if w.err == nil {
		w.message = emptyErrStr
	}
	return w.err
}

func (w *withErrors) GetMessage() string {
	return w.message
}

func (w *withErrors) GetPriority() string {
	if w.priority == "" {
		w.message = emptyPriorityStr
	}
	return w.priority
}

// Below are set process.
// All the process will be replace existing one if any.

func (w *withErrors) WithCode(c string) WithErrors {
	w.code = c
	return w
}

func (w *withErrors) WithCustomMessage(c interface{}) WithErrors {
	w.customMessage = c
	return w
}

func (w *withErrors) WithDependency(d string) WithErrors {
	w.dependency = d
	return w
}

func (w *withErrors) WithPriority(p Priority) WithErrors {
	var priority string
	switch p {
	case VERY_LOW:
		priority = veryLow
	case LOW:
		priority = low
	case MEDIUM:
		priority = medium
	case HIGH:
		priority = high
	case VERY_HIGH:
		priority = veryHigh
	}
	w.priority = priority
	return w
}
