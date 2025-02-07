package witherrors

import "errors"

type withErrors struct {
	code             string
	customMessage    interface{}
	dependency       string
	err              error
	forcedHttpStatus int
	message          string
	priority         string
}

type WithErrors interface {
	error
	// Get error code.
	Code() string
	// Get additional custom error message.
	CustomMessage() interface{}
	// Get error dependency.
	Dependency() string
	// Get error object.
	ErrorObject() error
	// Get additional error message.
	Message() string
	// Get error priority.
	Priority() string
	// Get forced error HTTP status.
	ForcedErrorHTTPStatus() int
	// Set error code.
	WithCode(c string) WithErrors
	// Set additional custom error message.
	WithCustomMessage(m interface{}) WithErrors
	// Set error dependency.
	WithDependency(d string) WithErrors
	// Set error priority.
	WithPriority(p Priority) WithErrors
	// Force error HTTP status.
	// This will be used for error handling in the API layer (only works for http).
	WithForcedErrorHTTPStatus(status int) WithErrors
}

// Create new error.
// The behaviour is same as go built in error.
func NewError(m string) WithErrors {
	if m == "" {
		return build(nil, emptyErrStr)
	}
	return build(nil, m)
}

// Wrap existing error to be passed into available struct attributes.
func WrapError(err error) WithErrors {
	if err == nil {
		return nil
	}

	return build(err, "")
}

func build(initialErr error, message string) WithErrors {
	if _, ok := initialErr.(*withErrors); !ok {
		if initialErr != nil {
			return &withErrors{
				err: initialErr,
			}
		}

		return &withErrors{
			err: errors.New(message),
		}
	}

	return initialErr.(*withErrors)
}

// Get the error message, implementing go built in Error() function.
func (w *withErrors) Error() string {
	return w.err.Error()
}

func GetCode(err error) string {
	var (
		wErr WithErrors
		ok   bool
	)

	if wErr, ok = err.(WithErrors); !ok {
		return ""
	}

	return wErr.Code()
}

func (w *withErrors) Code() string {
	if w.code == "" {
		w.message = emptyCodeStr
	}
	return w.code
}

func GetCustomMessage(err error) interface{} {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return nil
	}
	return wErr.CustomMessage()
}

func (w *withErrors) CustomMessage() interface{} {
	if w.customMessage == nil {
		w.message = emptyCustomMessageStr
	}
	return w.customMessage
}

func GetDependency(err error) string {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return ""
	}
	return wErr.Dependency()
}

func (w *withErrors) Dependency() string {
	if w.dependency == "" {
		w.message = emptyDependencyStr
	}
	return w.dependency
}

func GetErrorObject(err error) error {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return nil
	}
	return wErr.ErrorObject()
}

func (w *withErrors) ErrorObject() error {
	if w.err == nil {
		w.message = emptyErrStr
	}
	return w.err
}

func GetMessage(err error) string {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return ""
	}
	return wErr.Message()
}

func (w *withErrors) Message() string {
	return w.message
}

func GetPriority(err error) string {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return ""
	}
	return wErr.Priority()
}

func (w *withErrors) Priority() string {
	if w.priority == "" {
		w.message = emptyPriorityStr
	}
	return w.priority
}

func GetForcedHTTPStatus(err error) int {
	var (
		wErr WithErrors
		ok   bool
	)
	if wErr, ok = err.(WithErrors); !ok {
		return 0
	}
	return wErr.ForcedErrorHTTPStatus()
}

func (w *withErrors) ForcedErrorHTTPStatus() int {
	return w.forcedHttpStatus
}

// Below are set process.
// All of the process will replace the existing one if any.

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

func (w *withErrors) WithForcedErrorHTTPStatus(status int) WithErrors {
	w.forcedHttpStatus = status
	return w
}
