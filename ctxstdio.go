// Package ctxstdio provides overrideable access to stdio streams from a context.
package ctxstdio

import (
	"context"
	"io"
	"os"
)

type StdinType string
type StdOutType string
type StdErrType string

const StdinKey = StdinType("Stdin")
const StdOutKey = StdOutType("StdOut")
const StdErrKey = StdErrType("StdErr")

// Set alternate stdio streams
func Set(ctx context.Context, stdOut io.Writer, stdErr io.Writer, stdIn io.ReadCloser) context.Context {
	ctx = context.WithValue(ctx, StdinKey, stdIn)
	ctx = context.WithValue(ctx, StdOutKey, stdOut)
	ctx = context.WithValue(ctx, StdErrKey, stdErr)
	return ctx
}

// Stdin gets the stored stdin stream from the context
func Stdin(ctx context.Context) io.ReadCloser {
	if stdin, ok := ctx.Value(StdinKey).(io.ReadCloser); ok {
		return stdin
	}
	return os.Stdin
}

// StdOut gets the stored stdin stream from the context
func StdOut(ctx context.Context) io.Writer {
	if stdout, ok := ctx.Value(StdOutKey).(io.Writer); ok {
		return stdout
	}
	return os.Stdin
}

// StdErr gets the stored stdin stream from the context
func StdErr(ctx context.Context) io.Writer {
	if stderr, ok := ctx.Value(StdErrKey).(io.Writer); ok {
		return stderr
	}
	return os.Stdin
}
