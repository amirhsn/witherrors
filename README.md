# With Errors

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/amirhsn/witherrors/main/LICENSE)
[![Build Status](https://app.travis-ci.com/amirhsn/witherrors.svg?branch=main)](https://travis-ci.com/amirhsn/witherrors)
[![codecov](https://badgen.net/codecov/c/github/amirhsn/witherrors)](https://codecov.io/gh/amirhsn/witherrors)


With Errors (witherrors) package provide a flexible way to wrap and add additional information to your errors.

## Features
* [New Error](#new-error)
* [Wrap Error](#wrap-error)
* [Error](#error)
* [With Code](#with-code)
* [With Custom Message](#with-custom-message)
* [With Dependency](#with-dependency)
* [With Priority](#with-priority)
* [Get Code](#get-code)
* [Get Custom Message](#get-custom-message)
* [Get Dependency](#get-dependency)
* [Get Error](#get-error)
* [Get Message](#get-message)
* [Get Priority](#get-priority)
* [Clear](#clear)
* [Clear And Get](#clear-and-get)

## Installation
```bash
$ go get -u github.com/amirhsn/witherrors
```

## Usage
Below is the detail of how to use this module with listed features. Firstly we need to import this modules.
```go
import er "github.com/amirhsn/witherrors"
```

### New Error
Same as built in go error which is create new error that takes one string parameter.
```go
newErr := er.NewError("this is new error") 
```

### Wrap Error
Wrap existing error to be passed into `witherrors` struct attribute.
```go
_, err := redis.HDEL(.....)
return er.WrapError(err)
```

### Error
Get error message as string data type, this is implementation of built in `Error()` in Go.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err)
errMessage := myError.Error()
```

### With Code
Set error code as a string data type.
```go
_, err := redis.HDEL(.....)
return er.WrapError(err).WithCode("CODE-1")
```

### With Custom Message
Set error custom message as an empty interface data type.
```go
_, err := redis.HDEL(.....)
return er.WrapError(err).WithCode("CODE-1").WithCustomMessage("Custom Message")
```

### With Dependency
Set error dependency as a string data type.
```go
_, err := redis.HDEL(.....)
return er.WrapError(err).WithCode("CODE-1").WithCustomMessage("Custom Message").WithDependency("MONGO")
```

### With Priority
Set error priority as a `Priority` data type
```go
_, err := redis.HDEL(.....)
return er.WrapError(err).WithCode("CODE-1").WithPriority(er.HIGH)
```
This method will receive `Priority` data type which is consist of 5 values listed below.
* `VERY_LOW`
* `LOW`
* `MEDIUM`
* `HIGH`
* `VERY_HIGH`

### Get Code
Get error code as a string data type, if no code is set, `message` object will return empty code message.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err).WithCode("CODE-1")
code := myError.GetCode()
```

### Get Custom Message
Get custom message error as an empty interface data type, if no custom message is set, `message` object will return empty custom message.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err).WithCustomMessage("custom")
cMessage := myError.GetCustomMessage()
```

### Get Dependency
Get error dependency as a string data type, if no dependency is set, `message` object will return empty dependency message.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err).WithDependency("MONGO")
dep := myError.GetDependency()
```

### Get Error
Get error as an error data type, if no error is set, `message` object will return empty error message.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err)
newError := myError.GetError()
```

### Get Message
Get message as a string data type. Remember that this is not error message, but message because of empty get operation as mentioned above.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err)
message := myError.GetMessage()
```

### Get Priority
Get error priority as a string data type, if no priority is set, `message` object will return empty priority message.
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err).WithPriority(er.HIGH)
priority := er.GetPriority()
```

### Clear
Clear all existing data (code, dependency, priority, custom message).
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err)
myError.Clear()
```

### Clear And Get
Clear all existing data (code, dependency, priority, custom message).
```go
_, err := redis.HDEL(.....)
myError := er.WrapError(err)
clearedErr := myError.ClearAndGet().WithCustomMessage("new err")
```

## License
[MIT](https://github.com/amirhsn/witherrors/blob/main/LICENSE) © Amir Husein