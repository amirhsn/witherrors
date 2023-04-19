package witherrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	// Test empty error string
	err := NewError("")
	assert.EqualError(t, err, emptyErrStr)

	// Test non-empty error string
	err = NewError("test error")
	assert.EqualError(t, err, "test error")
}

func TestWrapError(t *testing.T) {
	// Test empty error string
	wrappedErr := WrapError(errors.New(""))
	assert.Equal(t, "", wrappedErr.Error())

	// Test non-empty error string
	inputErr := errors.New("test error")
	wrappedErr = WrapError(inputErr)
	assert.Equal(t, inputErr, wrappedErr.GetError())
}

func TestClear(t *testing.T) {
	w := &withErrors{
		code:          "code",
		customMessage: "customMessage",
		dependency:    "dependency",
		err:           errors.New("test error"),
		message:       "message",
		priority:      "priority",
	}
	w.Clear()

	assert.Equal(t, "", w.GetCode())
	assert.Equal(t, nil, w.GetCustomMessage())
	assert.Equal(t, "", w.GetDependency())
	assert.Nil(t, w.GetError())
	assert.Equal(t, "", w.GetPriority())
	assert.Equal(t, emptyPriorityStr, w.GetMessage())
}
func TestClearAndGet(t *testing.T) {
	w := &withErrors{
		code:          "code",
		customMessage: "customMessage",
		dependency:    "dependency",
		err:           errors.New("test error"),
		message:       "message",
		priority:      "priority",
	}
	w.ClearAndGet()

	assert.Equal(t, "", w.GetCode())
	assert.Equal(t, nil, w.GetCustomMessage())
	assert.Equal(t, "", w.GetDependency())
	assert.Nil(t, w.GetError())
	assert.Equal(t, "", w.GetPriority())
	assert.Equal(t, emptyPriorityStr, w.GetMessage())
}

func TestWithCode(t *testing.T) {
	w := &withErrors{}
	w.WithCode("test code")

	assert.Equal(t, "test code", w.GetCode())
}

func TestWithCustomMessage(t *testing.T) {
	w := &withErrors{}
	w.WithCustomMessage("test custom")

	assert.Equal(t, "test custom", w.GetCustomMessage())
}

func TestWithDependency(t *testing.T) {
	w := &withErrors{}
	w.WithDependency("test dependency")

	assert.Equal(t, "test dependency", w.GetDependency())
}

func TestWithPriority(t *testing.T) {
	w := &withErrors{}

	w.WithPriority(VERY_LOW)
	assert.Equal(t, veryLow, w.GetPriority())
	w.WithPriority(LOW)
	assert.Equal(t, low, w.GetPriority())
	w.WithPriority(MEDIUM)
	assert.Equal(t, medium, w.GetPriority())
	w.WithPriority(HIGH)
	assert.Equal(t, high, w.GetPriority())
	w.WithPriority(VERY_HIGH)
	assert.Equal(t, veryHigh, w.GetPriority())
}
