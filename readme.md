# Fiber Tools

This is a simple package with small tools to make fiber even nice then it already is.

## Install
```bash
$ go get github.com/incizzle/fibertools
```
```go
import fibertools "github.com/incizzle/fibertools"
```

## Usage

Error Handler. Set `x-debug` on request header to see stack trace.
```go
// Error handler that is passed into `fiber.New()`. Set x-debug in request header to see stack trace of error.
app := fiber.New(fiber.Config{
		ErrorHandler:          fibertools.ErrorHandler,
    })
```

Get Header
```go
// Pass in fiber context and the specific header you want to pull from the context.
fibertools.GetHeader(c, "user-agent")
```

New Error
```go
// Error function that can be returned on general errors || panics in order to get stack trace in error handler. Must be used in tandom with fibertools.Recover() for most value.
err, _ := errors.New("ERROR HERE")

return fibertools.NewError(err)
```

Recover
```go
// Simple fiber recover middleware based around default fiber recover but with the addiction of stack tarcing. Should be used with fibertools.ErrorHandler().
app := fiber.New(fiber.Config{
		ErrorHandler:          fibertools.ErrorHandler,
    })
    
app.Use(fibertools.Recover())
```

By: iNcizzle#1337
Twitter: [iNcizzle](https://twitter.com/incizzle)