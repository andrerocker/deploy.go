package http

import "io"

type Handler func(Request)

type Engine interface {
	Use(Filter)
	Start(string)
	Register(string, string, Handler)
}

type Request interface {
	Abort(int)
	Writer() io.Writer
	Reader() io.Reader
	ContextParameter(string) string
	RequestParameter(string) string
}

type Filter func(Request)
