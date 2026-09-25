// Package ctxstdio provides overrideable access to stdio streams from a context.
package ctxstdio

import (
	"context"
	"io"
	"os"
)

type stdinType string
type stdOutType string
type stdErrType string

const stdinKey = stdinType("Stdin")
const stdOutKey = stdOutType("StdOut")
const stdErrKey = stdErrType("StdErr")

// Set alternate stdio streams
func Set(ctx context.Context, stdOut io.Writer, stdErr io.Writer, stdIn io.ReadCloser) context.Context {
	ctx = context.WithValue(ctx, stdinKey, stdIn)
	ctx = context.WithValue(ctx, stdOutKey, stdOut)
	ctx = context.WithValue(ctx, stdErrKey, stdErr)
	return ctx
}

// Stdin gets the stored stdinsstream from the context
func Stdin(ctx context.Context) io.ReadCloser {
	if stdin, ok := ctx.Value(stdinKey).(io.ReadCloser); ok {
		return stdin
	}
	return os.Stdin
}

// StdOut gets the stored stdin stream from the context
func StdOut(ctx context.Context) io.Writer {
	if stdout, ok := ctx.Value(stdOutKey).(io.Writer); ok {
		return stdout
	}
	return os.Stdout
}

// StdErr gets the stored stdin stream from the context
func StdErr(ctx context.Context) io.Writer {
	if stderr, ok := ctx.Value(stdErrKey).(io.Writer); ok {
		return stderr
	}
	return os.Stderr
}
